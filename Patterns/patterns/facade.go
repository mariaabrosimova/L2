package patterns

import (
	"fmt"
	"os/user"

	"github.com/elastic/gosigar"
)

type UserInfo struct {
	username string
	name     string
	homeDir  string
}

func NewUserInfo() *UserInfo {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	return &UserInfo{
		username: usr.Username,
		name:     usr.Name,
		homeDir:  usr.HomeDir,
	}
}

func GetUserInfo(info UserInfo) {
	fmt.Println("Username: ", info.username)
	fmt.Println("Name: ", info.name)
	fmt.Println("HomeDir: ", info.homeDir)
}

type Memory struct {
	all  uint64
	used uint64
	free uint64
}

func NewMemory() *Memory {
	mem := gosigar.Mem{}
	mem.Get()
	return &Memory{
		all:  mem.Total / 1024,
		used: mem.Used / 1024,
		free: mem.Free / 1024,
	}

}
func GetMemory(info Memory) {
	fmt.Println("All memory: ", info.all)
	fmt.Println("Used memory: ", info.used)
	fmt.Println("Free memory: ", info.free)
}

type ComputerInfoFacade struct {
	userInfo *UserInfo
	memory   *Memory
}

func NewComputerInfoFacade() *ComputerInfoFacade {
	ComputerInfoFacade := &ComputerInfoFacade{
		userInfo: NewUserInfo(),
		memory:   NewMemory(),
	}
	return ComputerInfoFacade
}

func GetComputerInfoFacade(info ComputerInfoFacade) {
	GetUserInfo(*info.userInfo)
	GetMemory(*info.memory)
}

func Main_facade() {
	info := NewComputerInfoFacade()
	GetComputerInfoFacade(*info)
}
