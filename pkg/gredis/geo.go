package gredis

// Set a key/value
func GeoAdd(key string, longitude string, latitude string, address string) error {
	conn := RedisConn.Get()
	defer conn.Close()
	args := []interface{}{key, longitude, latitude, address}
	_, err := conn.Do("GEOADD", args...)
	if err != nil {
		return err
	}
	return nil
}
