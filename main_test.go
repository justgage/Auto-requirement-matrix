package main

func mainTest() {
	y := `

Reqirements: 
- req1
- req2
- req3


Design Entities: 
- name: Calendar
  description: This is a calendar entitiy
  reqirements:
   - req1
   - req2
   - req3

- name: User
  reqirements:
  - req3

- name: Calendar
  reqirements:
  - req1
  - req2
  - req3

- name: User
  reqirements:
  - req3
 `

}
