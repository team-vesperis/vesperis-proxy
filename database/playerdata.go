package database

import (
	"context"
	"database/sql"
	"encoding/json"
)

func SetPlayerData(playerId string, playerData map[string]interface{}) {
	redis := getRedisConnection()
	mysql, errConnection := getMySQLConnection(context.Background())
	if errConnection != nil {
		return
	}
	defer mysql.Close()

	// redis
	redisKey := "player:" + playerId
	for key, value := range playerData {
		jsonValue, _ := json.Marshal(value)

		err := redis.HSet(context.Background(), redisKey, key, jsonValue).Err()
		if err != nil {
			logger.Error("Error saving key and value to the Redis database. - ", err)
		}
	}

	// mysql
	jsonData, _ := json.Marshal(playerData)
	_, err := mysql.ExecContext(context.Background(), "REPLACE INTO player_data (playerId, data) VALUES (?, ?)", playerId, string(jsonData))
	if err != nil {
		logger.Error("Error saving player data to the MySQL database. - ", err)
	}
}

func SetPlayerDataField(playerId, field string, value interface{}) {
	redis := getRedisConnection()
	mysql, errConnection := getMySQLConnection(context.Background())
	if errConnection != nil {
		return
	}
	defer mysql.Close()

	// redis
	redisKey := "player:" + playerId
	jsonValue, _ := json.Marshal(value)
	err := redis.HSet(context.Background(), redisKey, field, jsonValue).Err()
	if err != nil {
		logger.Error("Error saving field to the Redis database. - ", err)
	}

	// mysql
	playerData := GetPlayerData(playerId)
	playerData[field] = value
	jsonData, _ := json.Marshal(playerData)
	_, err = mysql.ExecContext(context.Background(), "REPLACE INTO player_data (playerId, data) VALUES (?, ?)", playerId, string(jsonData))
	if err != nil {
		logger.Error("Error saving player data field to the MySQL database. - ", err)
	}
}

func GetPlayerDataField(playerId, field string) interface{} {
	redis := getRedisConnection()
	mysql, errConnection := getMySQLConnection(context.Background())
	if errConnection != nil {
		return nil
	}
	defer mysql.Close()

	// redis
	redisKey := "player:" + playerId
	value, err := redis.HGet(context.Background(), redisKey, field).Result()
	if err == nil {
		var result interface{}
		json.Unmarshal([]byte(value), &result)
		return result
	}

	// mysql
	var jsonData string
	err = mysql.QueryRowContext(context.Background(), "SELECT data FROM player_data WHERE playerId = ?", playerId).Scan(&jsonData)
	if err != nil {

		// no data found -> create new data
		if err == sql.ErrNoRows {
			emptyData := make(map[string]interface{})
			SetPlayerData(playerId, emptyData)
			return nil
		}

		logger.Error("Error getting player data field from the MySQL database. - ", err)
		return nil
	}

	// update redis
	var playerData map[string]interface{}
	json.Unmarshal([]byte(jsonData), &playerData)
	for key, value := range playerData {
		jsonValue, _ := json.Marshal(value)
		redis.HSet(context.Background(), redisKey, key, jsonValue)
	}

	return playerData[field]
}

func GetPlayerData(playerId string) map[string]interface{} {
	redis := getRedisConnection()
	mysql, errConnection := getMySQLConnection(context.Background())
	if errConnection != nil {
		return nil
	}
	defer mysql.Close()
	redisKey := "player:" + playerId

	// redis
	redisData, err := redis.HGetAll(context.Background(), redisKey).Result()
	if err == nil && len(redisData) > 0 {
		playerData := make(map[string]interface{})
		for key, value := range redisData {
			var fieldValue interface{}
			json.Unmarshal([]byte(value), &fieldValue)
			playerData[key] = fieldValue
		}
		return playerData
	}

	// mysql
	var jsonData string
	err = mysql.QueryRowContext(context.Background(), "SELECT data FROM player_data WHERE playerId = ?", playerId).Scan(&jsonData)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Warn("No player data found in MySQL for player: ", playerId)
			// Create new empty data
			emptyData := make(map[string]interface{})
			SetPlayerData(playerId, emptyData)
			return emptyData
		}
		logger.Error("Error getting player data from the MySQL database. - ", err)
		return nil
	}

	// update redis
	var playerData map[string]interface{}
	json.Unmarshal([]byte(jsonData), &playerData)
	for key, value := range playerData {
		jsonValue, _ := json.Marshal(value)
		redis.HSet(context.Background(), redisKey, key, jsonValue)
	}

	return playerData
}
