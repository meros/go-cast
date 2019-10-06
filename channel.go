package cast

import (
	"encoding/binary"
	"encoding/json"
	"log"
	"net"

	proto "github.com/golang/protobuf/proto"
)

type Channel struct {
	conn net.Conn
}

type ConnectionMessage struct {
	Type string `json:"type"`
}

func readJSONMessage() error {
	//TODO
	return nil
}

func sendJSONMessage(conn net.Conn, payload interface{}) error {
	// Prepare message
	protocolVersion := CastMessage_CASTV2_1_0
	sourceId := "sender-0"
	destinationId := "receiver-0"
	payloadType := CastMessage_STRING
	namespace := "urn:x-cast:com.google.cast.tp.connection"
	payloadData, err := json.Marshal(payload)
	message := &CastMessage{
		ProtocolVersion: &protocolVersion,
		SourceId:        &sourceId,
		DestinationId:   &destinationId,
		Namespace:       &namespace,
		PayloadType:     &payloadType,
		PayloadUtf8:     proto.String(string(payloadData)),
	}

	messageData, err := proto.Marshal(message)
	if err != nil {
		log.Println(err)
		return err
	}

	// Send frame header (4 byte B.E. length)
	var messageLen uint32 = uint32(len(messageData))
	err = binary.Write(conn, binary.BigEndian, &messageLen)
	if err != nil {
		log.Println(err)
		return err
	}

	// Send message
	_, err = conn.Write(messageData)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (channel *Channel) Connect() error {

	err := sendJSONMessage(channel.conn, &ConnectionMessage{"CONNECT"})
	if err != nil {
		log.Println(err)
		return err
	}

	err = sendJSONMessage(channel.conn, &ConnectionMessage{"PING"})
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

/*
go func(connection *tls.Conn) {
	for {

		var size uint32
		err := binary.Read(connection, binary.BigEndian, &size)
		if err == io.EOF {
			continue
		}

		if err != nil {
			log.Println(err)
			return
		}

		data := make([]byte, size)
		bytesRead, err := connection.Read(data)
		if err != nil || uint32(bytesRead) != size {
			log.Println(err)
			return
		}

		message := &CastMessage{}
		err = proto.Unmarshal(data, message)
		if err != nil {
			log.Println(err)
			return
		}

		log.Printf("Got message len %d %s", bytesRead, *message.PayloadUtf8)
		// TODO: Read message!
	}
}(connection)

protocolVersion := CastMessage_CASTV2_1_0
sourceId := "sender-0"
destinationId := "receiver-0"
payloadType := CastMessage_STRING

// Establish connection
namespace := "urn:x-cast:com.google.cast.tp.connection"
payload := "{ \"type\": \"CONNECT\" }"

message := &CastMessage{
	ProtocolVersion: &protocolVersion,
	SourceId:        &sourceId,
	DestinationId:   &destinationId,
	Namespace:       &namespace,
	PayloadType:     &payloadType,
	PayloadUtf8:     &payload,
}

data, err := proto.Marshal(message)
if err != nil {
	log.Println(err)
	return err
}

var size uint32 = uint32(len(data))
err = binary.Write(connection, binary.BigEndian, &size)
if err != nil {
	log.Println(err)
	return err
}

_, err = connection.Write(data)
if err != nil {
	log.Println(err)
	return err
}

namespace = "urn:x-cast:com.google.cast.tp.heartbeat"
payload = "{ \"type\": \"PING\" }"

message = &CastMessage{
	ProtocolVersion: &protocolVersion,
	SourceId:        &sourceId,
	DestinationId:   &destinationId,
	Namespace:       &namespace,
	PayloadType:     &payloadType,
	PayloadUtf8:     &payload,
}

data, err = proto.Marshal(message)
if err != nil {
	log.Println(err)
	return err
}

size = uint32(len(data))
log.Printf("writing message size %d", size)
err = binary.Write(connection, binary.BigEndian, &size)
if err != nil {
	log.Println(err)
	return err
}

_, err = connection.Write(data)
if err != nil {
	log.Println(err)
	return err
}

namespace = "urn:x-cast:com.google.cast.receiver"
payload = "{ \"type\": \"GET_STATUS\", \"requestId\" : 1 }"

message = &CastMessage{
	ProtocolVersion: &protocolVersion,
	SourceId:        &sourceId,
	DestinationId:   &destinationId,
	Namespace:       &namespace,
	PayloadType:     &payloadType,
	PayloadUtf8:     &payload,
}

data, err = proto.Marshal(message)
if err != nil {
	log.Println(err)
	return err
}

size = uint32(len(data))
log.Printf("writing message size %d", size)
err = binary.Write(connection, binary.BigEndian, &size)
if err != nil {
	log.Println(err)
	return err
}

_, err = connection.Write(data)
if err != nil {
	log.Println(err)
	return err
}

log.Printf("Here")

return nil*/
