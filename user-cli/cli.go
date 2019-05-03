package main

import (
	"log"
	"os"

	"github.com/micro/cli"

	pb "github.com/jglista/go-microservices/user-service/proto/user"
	"github.com/micro/go-micro"
	microclient "github.com/micro/go-micro/client"
	"golang.org/x/net/context"
)

func main() {
	srv := micro.NewService(
		micro.Name("go.micro.srv.user-cli"),
		micro.Version("latest"),
		micro.Flags(
			cli.StringFlag{
				Name:  "name",
				Usage: "Your name",
			},
			cli.StringFlag{
				Name:  "email",
				Usage: "Your email",
			},
			cli.StringFlag{
				Name:  "password",
				Usage: "Your password",
			},
			cli.StringFlag{
				Name:  "company",
				Usage: "Your company",
			},
		),
	)

	srv.Init(micro.Action(func(c *cli.Context) {
		// Check to ensure requried arguments were supplied to the cli
		if (c.String("name") == "") ||
			(c.String("email") == "") ||
			(c.String("password") == "") ||
			(c.String("company") == "") {
			log.Fatal("name, email, password, and company are required and were not passed in")
		}

		client := pb.NewUserService("go.micro.srv.user", microclient.DefaultClient)

		r, err := client.Create(context.TODO(), &pb.User{
			Name:     c.String("name"),
			Email:    c.String("email"),
			Password: c.String("password"),
			Company:  c.String("company"),
		})
		if err != nil {
			log.Fatalf("Could not create: %v", err)
		}
		log.Printf("Created: %s", r.User.Id)

		getAll, err := client.GetAll(context.Background(), &pb.Request{})
		if err != nil {
			log.Fatalf("Could not list users: %v", err)
		}
		for _, v := range getAll.Users {
			log.Println(v)
		}

		os.Exit(0)
	}))
}
