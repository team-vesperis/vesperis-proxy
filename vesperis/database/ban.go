package database

import (
	"context"
	"database/sql"
	"time"
)

func IsPlayerBanned(playerId string) bool {
	// redis
	redis := getRedisClient()
	redisKey := "banned_players:" + playerId

	value, err := redis.HGet(context.Background(), redisKey, "banned").Result()
	if err == nil {
		if value == "1" {
			return true
		} else if value == "0" {
			return false
		}
	}

	// mysql
	mysql, err := getMySQLConnection(context.Background())
	if err != nil {
		return true
	}
	defer mysql.Close()

	var exists bool
	err = mysql.QueryRowContext(context.Background(), "SELECT EXISTS(SELECT 1 FROM banned_players WHERE playerId = ?)", playerId).Scan(&exists)
	if err != nil {
		return true
	}

	// update redis
	redis.HSet(context.Background(), redisKey, "banned", exists)

	return exists
}

func BanPlayer(playerId, playerName, reason string) error {
	// mysql
	mysql, err := getMySQLConnection(context.Background())
	if err != nil {
		return err
	}
	defer mysql.Close()

	_, err = mysql.ExecContext(context.Background(), "REPLACE INTO banned_players (playerId, playerName, reason, permanently, ban_issued, ban_expires) VALUES (?, ?, ?, ?, ?, ?)", playerId, playerName, reason, true, time.Now(), sql.NullTime{})
	if err != nil {
		return err
	}

	// redis
	redis := getRedisClient()
	redisKey := "banned_players:" + playerId
	redis.HSet(context.Background(), redisKey, "banned", true)

	return nil
}

func UnBanPlayer(playerId string) error {
	// mysql
	mysql, err := getMySQLConnection(context.Background())
	if err != nil {
		return err
	}
	defer mysql.Close()

	_, err = mysql.ExecContext(context.Background(), "DELETE FROM banned_players WHERE playerId = ?", playerId)
	if err != nil {
		return err
	}

	// redis
	redis := getRedisClient()
	redisKey := "banned_players:" + playerId
	redis.HSet(context.Background(), redisKey, "banned", false)

	return nil
}

func TempBanPlayer(playerId, playerName, reason string, durationLength uint16, durationType time.Duration) error {
	// mysql
	mysql, err := getMySQLConnection(context.Background())
	if err != nil {
		return err
	}
	defer mysql.Close()

	expiration := time.Now().Add(time.Duration(durationLength) * durationType)

	_, err = mysql.ExecContext(context.Background(), "REPLACE INTO banned_players (playerId, playerName, reason, permanently, ban_issued, ban_expires) VALUES (?, ?, ?, ?, ?, ?)", playerId, playerName, reason, false, time.Now(), expiration)
	if err != nil {
		return err
	}

	// redis
	redis := getRedisClient()
	redisKey := "banned_players:" + playerId
	redis.HSet(context.Background(), redisKey, "banned", true)

	return nil
}

func CheckTempBans() {
	mysql, err := getMySQLConnection(context.Background())
	if err != nil {
		logger.Error("Error connecting to MySQL Database. - ", err)
		return
	}
	defer mysql.Close()

	rows, err := mysql.QueryContext(context.Background(), "SELECT playerId, ban_expires FROM banned_players WHERE permanently = ?", false)
	if err != nil {
		logger.Error("Error querying banned players. - ", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var playerId string
		var banExpires sql.NullTime
		err := rows.Scan(&playerId, &banExpires)
		if err != nil {
			logger.Error("Error scanning banned player data. - ", err)
			continue
		}

		if banExpires.Valid && time.Now().After(banExpires.Time) {
			logger.Info("Player with ID: " + playerId + " has been unbanned due to expiring temporarily ban.")
			UnBanPlayer(playerId)
		}
	}
}

func GetBanReason(playerId string) string {
	mysql, err := getMySQLConnection(context.Background())
	if err != nil {
		return ""
	}
	defer mysql.Close()

	var reason string
	err = mysql.QueryRowContext(context.Background(), "SELECT reason FROM banned_players WHERE playerId = ?", playerId).Scan(&reason)
	if err != nil {
		return ""
	}

	return reason
}

func GetBanExpiration(playerId string) time.Time {
	mysql, err := getMySQLConnection(context.Background())
	if err != nil {
		return time.Now().Add(time.Hour)
	}
	defer mysql.Close()

	var banExpires sql.NullTime
	mysql.QueryRowContext(context.Background(), "SELECT ban_expires FROM banned_players WHERE playerId = ?", playerId).Scan(&banExpires)

	if banExpires.Valid {
		return banExpires.Time
	}

	return time.Now().Add(time.Hour)
}

func IsPlayerPermanentlyBanned(playerId string) bool {
	if !IsPlayerBanned(playerId) {
		return false
	}

	mysql, err := getMySQLConnection(context.Background())
	if err != nil {
		return true
	}
	defer mysql.Close()

	var permanently bool
	err = mysql.QueryRowContext(context.Background(), "SELECT permanently FROM banned_players WHERE playerId = ?", playerId).Scan(&permanently)
	if err != nil {
		return true
	}

	return permanently
}

func GetBannedPlayerNameList() []string {
	var list []string
	mysql, err := getMySQLConnection(context.Background())
	if err != nil {
		return list
	}
	defer mysql.Close()

	rows, err := mysql.QueryContext(context.Background(), "SELECT playerName FROM banned_players")
	if err != nil {
		logger.Error("Error querying banned players. - ", err)
		return list
	}
	defer rows.Close()

	for rows.Next() {
		var playerName string
		err := rows.Scan(&playerName)
		if err != nil {
			logger.Error("Error scanning banned player data. - ", err)
			continue
		}
		list = append(list, playerName)
	}

	return list
}

func GetBannedPlayerIdByName(playerName string) string {
	mysql, err := getMySQLConnection(context.Background())
	if err != nil {
		return ""
	}
	defer mysql.Close()

	var playerId string
	err = mysql.QueryRowContext(context.Background(), "SELECT playerId FROM banned_players WHERE playerName = ?", playerName).Scan(&playerId)
	if err != nil {
		return ""
	}

	return playerId
}
