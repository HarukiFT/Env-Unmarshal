
# Env-Unmarshal

Test module for practicing reflection in Go

## Main structure - Env
**LoadEnv(path string)** returns new Env structure's object, where path it's path relate work dir

**ToEnv()** write env data to process env

**Unmarshal(target interface{})** unmarshal env data to structure object
