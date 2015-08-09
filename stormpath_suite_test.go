package stormpath_test

import (
	"log"
	"os"
	"runtime"
	"testing"

	"github.com/garyburd/redigo/redis"
	. "github.com/jarias/stormpath-sdk-go"
	uuid "github.com/nu7hatch/gouuid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	app     *Application
	cred    Credentials
	account *Account
	tenant  *Tenant
)

const (
	string256  = "5mmsUvVLGAfBphgbkWoNjsPVLafEUaqJFKSKbHJkDsuWFOUTq2q3SRQSbSnT7J9hGyDCBtsACTTgreSKIRtRci1lywp0g2J474tmnfCEHaSItpFFtWvkOr4IJgmZ0ZYacnzWF4JJjC6KZAOmAOIG0AwUxJ22TTEnkBToWrzvTEunuwXBZfWcyQKToEsj1QCeHNJ7OOfQYClELLxjAvSVEQSRYJta72LZICRPsoTl4aMYRbgD92l83vmCCHma4qOk"
	string4001 = "a9j5ko04xuMRmj2mP2Ex2Ue7DtZ62UDPKxPtQvRvqPxVNPxsylszp0RIpyMDeKuMB7AMoU4mWbrjEsVR5JGe64ZqzuHlEKcBSJ8Ci5H2LsZ1Le2JbDWKIQjGLTHaKgC0Ecu3Xe8RpMLo72YbKYESQ2YBzOKczDphPi7iovGrxGx5uoyw0Zh5XfBMTLZlXmSfpzB2jhNpyCHBUhlFittJ21Ff8GSV5eTE8EtnjpwvseirWBQjwOa9I6sqfQUMuKJwTqt69vtZa0YGN81aDtylqF1vqIjWLp9WmNnpD5r17xHez0nmjT0I7zaMGsyvmSXpPHyLipLw9kKspHDpEZTXYbX6pnMosjjyCE9Nina4xBsp2ycMzpKyMllaRFFI8t1zB5mcTTfVTxwaBDO3JS66ASJONVtLSlyWVWP4c8guLEXcZQczLQwDr0Z30ukXxlYJ6Wq5CjUcZKg45h1PEq3BYfeBmztrzg9kVaW8UfI56vYsutAYDeEuaoYCALZ5k1JOE4KFViaCak1jF29DqA0u0uYRHsNMKE9UPIXRHt3SsTrA7vSD2SEK7h77lV2tluengZXbaV0Sv8ZlanOrvTGgAj2REQkrqALm5XcOfcWCkYMgnJnB2VgCSyP5ORQR3CW1IS4o2ueCagqgrsPSE1goOEnL62ZQYFJU3mhHNRuxa9Hjo3I3XpDvwAMtN7nc55NLOceT9pQmghqEXFPqiwA1lTfkvTAXhxX05oJ7LseFknY8LW6F3C7B5a6yeDb2snV8DJQJJ7IW8gfamBsfMhcoBTw69KjnkRgT21bc7zpk5SpjuqPmaOzPWrmmMc5AQYx7GFuUlmfULmlxUWElG4JVhf7njBO86v7SjQ3WlV0LS9AbnEr4Myg5QZSa5J6Uq5CMQY8pVlybiZUZjh6JFcHGN2b8WLCo3LJf0qYq1G5pkkW3ehZyVJEU80P0rLvXv8BO73C0HguqaNbcij37j19a3Y4r7z6rwItkHX8XZysvUiIgunYRLHUMGDkUYDLb0AO8iEHz5OlM7tGOGL8il5pt9b7gJPBP8shvrLwhF1wEyXREOK8y6nHNoe7At3AhEmKIGc1tMiCrXRK1sHhQq5W5tcksfBulK4VWmU3OkxsJFhpLLNJoCYTx4ovXs2hKMtNAVkE8sU6TAF8fDAxRinF5kHrnXpUAQvm1RkaI8f4Coaze77EON3qenxGoFVUiIzl6l3xOby4sPP2XCJ2EDXEYutzePb916SBB5R8BfoGN85nVYkMiY7TCHbRNmPv1riQK1OYbMtSwZ6B4wkgiB0DJNEjEB54NRGT2ZeImCaL4oLS6xsyha1RKusXlfOy7Zw9KENh4R8uHwRxYq6KkewLBTbHW2j2cbGAvmhHXk3MD1MIV94jvNfn0rjmzw1TDigBrvTjLHgy46lV92cxOaRyW8OO2EDNRmbKrnu6WmxyYNpGaFG0H8xqpQNm10uIgKwPzza6O0qhgCGB0sQlg5WmD294ePNCSiURFb0VXkq02HqDLgyrbqNAArIFyN5BE07orRG7RRroaZZxkAmFTtj9r1O0zHB089jsSbLuV9LtOHLVlcU1p6JWZFY81TBe87oKU10Fj59rP5WekXkPwglS0begn2cyCOMX8bMnu2v0vqIB0IV3KCCmw4Oc4ieesfm6HS0861mIF9RcyPs6LhNDiuv2iRykmQqVymO3rx7HGFYX9ESqsQv3xawlPYGn6wayS29kDLnWHO9kElElfpTDzavJGs8RIYlWwbuq4PNcIuvTNcqNCFLuxuRXP0uTA7Naqrcc8BswiupLiuT2v4suu5gTW3PZy9UIKkx11jsuhqiCc5DyB1wlwrNVZJ92QQbki0xLoaF8CNC5B0mr4OChQnBjHBPnTsCUq5poDIeLEhKuiWDl8n0LOHzy2KVxNsvh7cavBBvsVYtuDuv3ZFuAt84TMmoUblQbQjOBBk9HPcCLwX42S16VXOZkJnvMQj80lAWFVbeEI0evrpVNh4qmDFFBqr8s2e1QmNT2HqZVEOpYATLgHjWqfgAKrXLeNh5NE6t85rSPMHxvSn4g6hgiW0nOCrVibNfReBRGx4GaTv6WfLVU5Rvwn4RmQpcTT4YbtXHWDpS7Y4Xi7bCPzkMrmC6LIwuDhQvVRf003xZ6jHwRPAzWOCfuOOofS2YGOX4wtOjuqE1SOpZoPMovH9LaAiLTKEEl8jNhisQDOnHZK0sl3j5r0Xh9o0QM2A1AC3L1nig28vylAm1HPxjDinMqgjbRWLobLPAPovt48Gw5VIGNbSLj4PoNoY3z2yQ6PBPr3KFVAf60lqoapRp3QWYnrjckT3cNuWXn0pgTMkTSJ68WABJvDR74j6ivTryQqDwluUp7yRumG93yqeY4xTOXemi9KHDL4UqbUk00RCD7y7SavuIt3F3HqUegSPBY30xjkeP3VGhAWISJaW3OO5qR5g9v2n7LU9OhP31AVbzC1Omj3JNgZ7sR4SSC11jrUMUeXMrxv3A8ALNhTE0IxZc0EUe4kCEYvbgZv06pPeAOSQbqtsfWTBtj37a9bltwOTBBul1TljrFqmGBYE20WphHpBwQQG8LW2WvR1wRcbfWK0AMYby0JxCMbobe51TZQEvvweuMuhbkuDoqEiz1QkZZUZRSg4woNkm0FGZoOuK5shKVESzEYANJexyNCKaMnXApBwZ9ohlTxETh92YXz9NqRuZYKfQWQHi05Bb6OoI2H8SzxE2wJkh0CEDgutlPcwhbUN00KLBaBNWj6ClGGHtmoZy3RK9EfKI3RKg84fqMoFRHuA1IOIeK3FuEefaUaSX2oExWqkT0sqasko9mfMNsxp0bLWnQ28BDHzi2VCJE5tbPKVB5hmbNOYwm3LVJmm017JuX2QRNLzr8B1jb0kmQrEr1ZsMe5YwzC5BYouZ5TDntNmtJOCCVSEV0xFbfThHiR1DmG4S2YbzW8UnTtTXwCnpyD0va3tk7n8QOu6WDJAuTMcvt5hNcghR5z4GnlNwpUrAOh2Ls6GA9vXCifWQBXuAWSJ9QomrQ0X0eBU0l6ri19fyTS5XQS56o4AVxQGfmcckR6Ii0THHHjv30ND0t1TlRyTE8nHHjhIz3aGFtGfvyxKNA22zev6AHKefJVr1QOofWDMKRkEoDQMqUCE2r8iJlUk4ZEUiBaJtk7KW4GHZxq8pNMywUVl9rukgppQIVDRTe3AuE7Nvef0xCUnHN9mEQs8t2TtLm6Y8stziorzr9MnHQEgmo7qoD5GhR6BTfCvz0ZHNFQDHjjIfi4K0Dsyo9rgelhKBQ4csyZDnFCD401unltlOTsHGTKFVjX6Sx3oIYRWYfSh2w4NfDT1LU6AchvE4t9eSCvwbrY7gKEN8HNHeRWVpvDWx3lNvGWZM6m0n1xwQ4xYu2YwzlUKKowJ2e2nJ3xRZ6MAUOp05YoRMGzI2Zp09R8B0K13p4s8DHy75V9TLk5foMT16MX1MGScjru9qZcnv5l4SyfKJzN8JsNnhT9jYis5lRCGUG5aGcYCLOUcQYDF5zbfYTsUfRl9XjRTA4GBz5luxp5qgzODPsSh4TuDusiFVYrBZTIpu0vrhpBXm8e1s9amOPZ37xmJyMvEqSaUt2UHFVhR0Mfz16Bxkxs9juMmYF4hnRpCQ5ewQf00xj2uEIlSjaGOYwxLMF1A4869jR6W2x7x1nLUxHM7xBAKthyBQw63bL9iH4MvUTgfhMjal4slmFB01lFxBJU5B7GaxOhf1Y2WQRKUej9PUeTXsy0GBPjSuApfuUbXkMKYXIUwBtXP7eXTQ5pkSaEpmmq0q8qb36Zo4B9cWHeHeVet7uOwDIj4SnPGAR9x4PJEx2DNtVLVix8OGtu9FczEgJCQLnXQeIHI4cUvKW4CecDz2wg2222kTkXvBUcImtJFkzvaOPS4c1jbnFQqF2MDxPzGemG7gA6Xwqsycl3u1j1fLvYRcMI26BO1YH2Y4VBu9uvGOIRWvpov2mXVDHy7bYHMQoKD"
	string1001 = "aLZrsWvzwK7amj0atmPfP1HLz3zeLzXuIh6TbkgxauexrsuDEUI40M86R9H8pTNekWRunKlUnjhH5ZRlRB9EMQlXHIho2ZFeZoPvTXz9Evm0gSI0qZabJtAizSmFAymCu6oUCwgKyaNQj1wqqfh8IyYzZNMY8njEUXRrRHkAVID87Fs0VeuApO3Ei6GPZ7EKZ0UnYzRiTjtP66cjYzYGuj4BKXe5MHKPe38vXDAFwWGHvIHGj8KzJC1z5NiPTUNMH6GOQKFVmw8NS6FpPx0yBRm0cbtUe9nuZBiMS76baZDvQIsNDvLyJGfXzOc0Dqm20RGiQhI1Da9JVehF60Ug5BDpnFKGzwRXBkkvLLNMsoKbE5H6w19IzOzKUgJOTkT59mbZUb6uEfAI6fKNMUFCtgosi3aM43xmhcz06vEOv1jRfXil12AnHXSOSLfupYzw0T0z1ywvuNhV1GGEXjUYIERxwebUUXLkHlWyzwsuRf2EF0umKeSQDH3vN43rXjKsv2ZB6JYJbtjebwPJGZ8M6YsaTIpyiksH5cB6mzOWbuaEtAqOj9FPjiI4KWrkiGBjbKDFLSKih5f6wwmTNj4knifB7VbIE4f2kPhcK4SCPiA6ifQpohLycqMAMIj1AfT6bPnyQQMkcEerVFbXbLgQhS4kRYwbY87huNFI7aJsyuhOcz9g6hDK3Z3b5A1BmND8qfpPrlAnmaPtWN0hNxecDQbBUFNBKotvF1WQ49YYL7oPprO5WDp40Z0PAp7GPyfTx6NCWl99fsHlqFHHD8FToW7ZsIG0n4SfQO67GnT93ir8VIjeGlLLwGu5Hrcb1j3RMlOrWFw3bO7mVfV45aY2SFeWWwuZurUm5JyTEVG3WES7FaO1rRsw0H9jlxWE1Ey1hJivMEtounUMW2bnDsYPHEyNoWxD6GwSCelHVJioSvIcgU3IFxbN9AtFFXabWU7qqcCIx6L3knzmgM0ByUse4sR0F2e09xGT757zhayvP"
)

func TestStormpath(t *testing.T) {
	runtime.GOMAXPROCS(4)
	RegisterFailHandler(Fail)
	RunSpecs(t, "Stormpath Suite")
}

func randomName() string {
	uuid, _ := uuid.NewV4()
	return uuid.String()
}

func newTestApplication() *Application {
	return NewApplication("app-" + randomName())
}

func newTestGroup() *Group {
	return NewGroup("group-" + randomName())
}

func newTestDirectory() *Directory {
	return NewDirectory("directory-" + randomName())
}

func newTestAccount() *Account {
	name := randomName()
	email := name + "@test.org"
	return NewAccount(email, "1234567z!A89", email, "givenName", "surname")
}

func initLogInTestMode() {
	Logger = log.New(GinkgoWriter, "", log.Ldate|log.Ltime|log.Lshortfile)
}

var _ = BeforeSuite(func() {
	var err error
	cred, err = NewDefaultCredentials()
	if err != nil {
		panic(err)
	}

	stormpathBaseURL := os.Getenv("STORMPATH_BASE_URL")
	if stormpathBaseURL != "" {
		BaseURL = stormpathBaseURL
	}

	cacheEnabled := os.Getenv("CACHE_ENABLED")
	if cacheEnabled == "true" {
		redisServer := os.Getenv("REDIS_SERVER")
		redisConn, err := redis.Dial("tcp", redisServer+":6379")
		if err != nil {
			panic(err)
		}

		Init(cred, RedisCache{redisConn})
	} else {
		Init(cred, nil)
	}
	initLogInTestMode()

	tenant, err = CurrentTenant()
	if err != nil {
		panic(err)
	}

	app = newTestApplication()

	err = tenant.CreateApplication(app)
	if err != nil {
		panic(err)
	}

	account = newTestAccount()
	account.Email = "test@test.org"
	account.Username = "test@test.org"
	app.RegisterAccount(account)
})

var _ = AfterSuite(func() {
	if app != nil {
		app.Purge()
	}
})
