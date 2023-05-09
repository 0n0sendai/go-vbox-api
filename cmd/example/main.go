package main

import (
	"log"
	"os"

	"github.com/0n0sendai/go-vbox-api/soap"
	"github.com/0n0sendai/go-vbox-api/vbox"
)

func logon(service vbox.VboxPortType, user string, pass string) (string, error) {
	session, err := service.IWebsessionManager_logon(&vbox.IWebsessionManager_logon{Username: user, Password: pass})

	// any error?
	if err != nil {
		return "", err
	}

	return session.Returnval, nil
}

func logoff(service vbox.VboxPortType, session string) error {
	_, err := service.IWebsessionManager_logoff(&vbox.IWebsessionManager_logoff{RefIVirtualBox: session})

	// any error?
	if err != nil {
		return err
	}

	return nil
}

func main() {
	url := "http://localhost:18083/"
	user := os.Getenv("VBOX_USER")
	pass := os.Getenv("VBOX_PASS")

	var session string = ""

	// create SOAP client
	client := soap.NewClient(url)
	// init service
	service := vbox.NewVboxPortType(client)

	// get a session
	session, err := logon(service, user, pass)
	// any error?
	if err != nil {
		// sorry
		log.Fatalf("ERROR: %v", err)
	}

	// print session
	log.Println("Logon: ", session)

	// get api version
	apiVersion, err := service.IVirtualBox_getAPIVersion(&vbox.IVirtualBox_getAPIVersion{This: session})
	// any error?
	if err != nil {
		// logoff
		if err := logoff(service, session); err != nil {
			// sorry
			log.Fatalf("ERROR: %v", err)
		}

		// sorry
		log.Fatalf("ERROR: %v", err)
	}

	// print API version
	log.Println("APIVersion: ", apiVersion)

	// create machine
	machineID, err := service.IVirtualBox_createMachine(
		&vbox.IVirtualBox_createMachine{
			This: session,
			Name: "test_machine",
		})
	// any error?
	if err != nil {
		// logoff
		if err := logoff(service, session); err != nil {
			// sorry
			log.Fatalf("ERROR: %v", err)
		}

		// sorry
		log.Fatalf("ERROR: %v", err)
	}

	// print machine ID
	log.Println("MachineID: ", machineID)

	// register machine
	_, err = service.IVirtualBox_registerMachine(
		&vbox.IVirtualBox_registerMachine{
			This:    session,
			Machine: machineID.Returnval,
		})
	// any error?
	if err != nil {
		// logoff
		if err := logoff(service, session); err != nil {
			// sorry
			log.Fatalf("ERROR: %v", err)
		}

		// sorry
		log.Fatalf("ERROR: %v", err)
	}

	// get session to lock for editing
	sessionObject, err := service.IWebsessionManager_getSessionObject(
		&vbox.IWebsessionManager_getSessionObject{
			RefIVirtualBox: session,
		})
	// any error?
	if err != nil {
		// logoff
		if err := logoff(service, session); err != nil {
			// sorry
			log.Fatalf("ERROR: %v", err)
		}

		// sorry
		log.Fatalf("ERROR: %v", err)
	}

	// lock machine
	lockTypeWrite := vbox.LockTypeWrite

	_, err = service.IMachine_lockMachine(
		&vbox.IMachine_lockMachine{
			This:     machineID.Returnval,
			Session:  sessionObject.Returnval,
			LockType: &lockTypeWrite,
		})
	// any error?
	if err != nil {
		// logoff
		if err := logoff(service, session); err != nil {
			// sorry
			log.Fatalf("ERROR: %v", err)
		}

		// sorry
		log.Fatalf("ERROR: %v", err)
	}

	// get machine for writing
	writeMachineID, err := service.ISession_getMachine(
		&vbox.ISession_getMachine{
			This: sessionObject.Returnval,
		})
	// any error?
	if err != nil {
		// logoff
		if err := logoff(service, session); err != nil {
			// sorry
			log.Fatalf("ERROR: %v", err)
		}

		// sorry
		log.Fatalf("ERROR: %v", err)
	}

	// edit machine
	_, err = service.IMachine_setMemorySize(
		&vbox.IMachine_setMemorySize{
			This:       writeMachineID.Returnval,
			MemorySize: 1024,
		})
	// any error?
	if err != nil {
		// logoff
		if err := logoff(service, session); err != nil {
			// sorry
			log.Fatalf("ERROR: %v", err)
		}

		// sorry
		log.Fatalf("ERROR: %v", err)
	}

	// save machine
	_, err = service.IMachine_saveSettings(
		&vbox.IMachine_saveSettings{
			This: writeMachineID.Returnval,
		})
	// any error?
	if err != nil {
		// logoff
		if err := logoff(service, session); err != nil {
			// sorry
			log.Fatalf("ERROR: %v", err)
		}

		// sorry
		log.Fatalf("ERROR: %v", err)
	}

	// unlock machine
	_, err = service.ISession_unlockMachine(
		&vbox.ISession_unlockMachine{
			This: sessionObject.Returnval,
		})
	// any error?
	if err != nil {
		// logoff
		if err := logoff(service, session); err != nil {
			// sorry
			log.Fatalf("ERROR: %v", err)
		}

		// sorry
		log.Fatalf("ERROR: %v", err)
	}

	// get machines
	machines, err := service.IVirtualBox_getMachines(
		&vbox.IVirtualBox_getMachines{
			This: session,
		})
	// any error?
	if err != nil {
		// logoff
		if err := logoff(service, session); err != nil {
			// sorry
			log.Fatalf("ERROR: %v", err)
		}

		// sorry
		log.Fatalf("ERROR: %v", err)
	}

	log.Println("Machines: ", machines.Returnval)

	// logoff
	if err := logoff(service, session); err != nil {
		// sorry
		log.Fatalf("ERROR: %v", err)
	}
}
