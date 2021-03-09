package main

import (
	bankgrpcv1 "agohomework6/pkg/bank/v1"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"os"
	"time"
)

const defaultPort = "9999"
const defaultHost = "0.0.0.0"

func main() {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = defaultPort
	}

	host, ok := os.LookupEnv("HOST")
	if !ok {
		host = defaultHost
	}

	if err := execute(net.JoinHostPort(host, port)); err != nil {
		log.Println(err)
		os.Exit(1)
	}

}

func execute(addr string) (err error) {

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer func() {
		if cerr := conn.Close(); cerr != nil {
			if err == nil {
				err = cerr
				return
			}
			log.Println(err)
		}
	}()

	client := bankgrpcv1.NewTemplateServiceClient(conn)
	//ctx, _ := context.WithTimeout(context.Background(), time.Second)
	//
	//var newtemplate = bankgrpcv1.MakeTemplate{
	//	Name:  "testtemplate",
	//	Phone: "1234567",
	//}
	//id, err := client.CreateTemplate(ctx, &newtemplate)
	//
	//if err != nil {
	//	if st, ok := status.FromError(err); ok {
	//		log.Println(st.Code())
	//		log.Println(st.Message())
	//	}
	//	return err
	//}
	//
	//log.Printf("template %d was created", id.Id)

	var all = bankgrpcv1.All{}
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	templatesList, err := client.GetAllTemplates(ctx, &all)
	if err != nil {
		if st, ok := status.FromError(err); ok {
			log.Println(st.Code())
			log.Println(st.Message())
		}
		return err
	}

	//log.Println(templatesList)

	var template = bankgrpcv1.TemplateId{Id: 1}
	ctx, _ = context.WithTimeout(context.Background(), time.Second*100)
	templatesList, err = client.GetTemplateById(ctx, &template)
	if err != nil {
		if st, ok := status.FromError(err); ok {
			log.Println(st.Code())
			log.Println(st.Message())
		}
		return err
	}

	log.Println(templatesList)

	return nil
}
