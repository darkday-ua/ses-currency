package db

import (
	"bufio"
	"os"
	"sync"
	"log"
	"fmt"
	"currency_service/config"
)

var	packageVersion string = "0.0.2"

type DB struct {
	file         *os.File
	buffer       []string
	cache        map[string]bool
	cacheLock    sync.RWMutex
	filePath     string
	inMemoryOnly bool
}

type SessionManager struct {
	db *DB
}

var sessionManager *SessionManager
var once sync.Once

func GetSessionManager() (*SessionManager, error) {
	once.Do(func() {
		db := &DB{
			filePath: config.Config.DB_PATH,
			cache:    make(map[string]bool),
		}

		// It would be better to have real database or even have memcache
		// because seeking in the file is quite expensive
		// So we emulate here serious things with write-through memcaching and db sessions :)

		err := db.fillMemoryCacheFromFile()
		if err != nil {
			log.Println("Can not populate memcache from the db file ", err)
			
		}else{
			fmt.Println("Populated memcache from the db file, ", db.cache)
		}

		sessionManager = &SessionManager{db: db}
	})

	return sessionManager, nil
}

func (sm *SessionManager) GetDBSession() *DB {
	return sm.db
}

func (db *DB) fillMemoryCacheFromFile() error {
	if db.filePath == "" {
		log.Println("filePath is not set")
		return nil
	}

	file, err := os.OpenFile(db.filePath,os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		db.cache[line] = true
	}

	return scanner.Err()
}

// #ATTENTION Should be better have these operation in separated CRUD module
func (db *DB) AddUser(user string) (string, bool) {
	db.cacheLock.RLock()
	_, exists := db.cache[user]
	db.cacheLock.RUnlock()

	if exists {
		return "user_exists", false
	}

	if db.filePath != "" {
		if file, err := os.OpenFile(db.filePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend); err == nil{	
			file.WriteString(user + "\n"); 
			file.Close()
		}
	}

	db.cacheLock.Lock()
	db.cache[user] = true
	db.cacheLock.Unlock()

	return "user_added",true
}

func (db *DB)GetSubscribedUsers() []string {
	db.cacheLock.RLock()
	users := make([]string, len(db.cache))

	i := 0
	for k := range db.cache {
		users[i] = k
		i++
	}
	db.cacheLock.RUnlock()	

	return users
}


func init() {	
	fmt.Println("..db package version ", packageVersion)
}

