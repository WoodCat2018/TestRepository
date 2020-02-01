package main

import (
	"fmt"
	"time"

	"github.com/Altarrel/goroyale"
)

var token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTIyMywiaWRlbiI6IjQ2OTA5NTQ0MzE1OTE4NzQ4NyIsIm1kIjp7fSwidHMiOjE1MzE5MTE3MjI0MDB9.uXafwrzwEjgJ7Y-HJEPlEALumU-Ts8b6Sha5J2AoIqs"

// Pair test struct
type Pair struct {
	Key   string
	Value int
}

// PairList test type
type PairList []Pair

func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value > p[j].Value }

func mapsort() {
	// MemberData := make(PairList, len(clan.Members))
	// sort.Sort(MemberData)
	// fmt.Println(MemberData)
	// MemberData[i] = Pair{clanmember.Name, clanmember.Donations}
}

func main() {
	c, err := goroyale.New(token, 0) // 0 will use the default request timeout of 10 seconds
	if err != nil {
		fmt.Println(err)
		return
	}

	ver, err := c.APIVersion()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("API Version:", ver)
	}

	params := map[string][]string{
		"exclude": {"name"},
	}
	clan, err := c.Clan("8CU9LQJ", params)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Len:", len(clan.Members))
		NameDate := make(map[string][]string, len(clan.Members))
		RoleDate := make(map[string][]string, len(clan.Members))
		DonationsData := make(map[string][]int, len(clan.Members))
		ReceivedData := make(map[string][]int, len(clan.Members))
		for _, clanmember := range clan.Members {
			fmt.Println("Name:", clanmember.Name, "Role", clanmember.Role, "Donations:", clanmember.Donations, "Received:", clanmember.DonationsReceived)

		}

	}

	clanwarlog, err := c.ClanWarLog("8CU9LQJ", params)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Len:", len(clanwarlog))
		for logN, logdata := range clanwarlog {
			nowtime := time.Now().Unix() - 604800
			if int64(logdata.CreatedDate) < nowtime {
				return
			}
			fmt.Println("NowTime:", nowtime)
			fmt.Println("No.", logN, "	Len:", len(logdata.Participants))
			fmt.Println("CreatedDate:", logdata.CreatedDate)
			for _, MemberDate := range logdata.Participants {
				fmt.Println("TAG:", MemberDate.Tag, "	NAME:", MemberDate.Name, "	Card", MemberDate.CardsEarned, "Win", MemberDate.Wins)
			}
		}
	}

}
