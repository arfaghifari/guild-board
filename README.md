# guild-board

## IDEAS
Do u need a hand to help but dont know how to ask? Now There is a GUILD BOARD. U can request a quest to guild, and our adventurer will go to help u. Dont forget to bring the commission after the quest have completed

## List API
### GET /quest-status  ~ ~ Get All Quest
Query : "status" = 0|1 

Body : {}

Example Response

### POST /quest  ~ ~ Make a quest
Request Body
```json
 {
    "name": "menyelamatkan kucing",
    "description" : "menyelamatkan kucing tersangkut di pohon",
    "minimum_rank" : 11,
    "reward_number" : 200000

}
```

Response

```json
success
```

### DELETE /quest  ~ ~ Delete a quest
Request Body
```json
 {
    "quest_id": 1

}
```

Response

```json
success
```

### PATCH /quest-rank  ~ ~ Update rank quest
Request Body
```json
 {
    "quest_id": 1,
    "minimum_rank" : 12

}
```

Response

```json
success
```

### PATCH /quest-rank  ~ ~ Update rank quest
Request Body
```json
 {
    "quest_id": 1,
    "reward_number" : 250000

}
```

Response

```json
success
```

### POST /adventurer  ~ ~ Make an adventurer

Request Body
```json
 {
    "name": "andi",
    "rank" : 11

}
```

Response

```json
success
```

### PATCH /adventurer-rank  ~ ~ Update rank an adventurer

Request Body
```json
 {
    "name": "andi",
    "rank" : 11

}
```

Response

```json
success
```

### POST /take-quest  ~ ~ An Adventurer take a quest

Request Body
```json
 {
    "adv_id": 1,
    "quest_id" : 1

}
```

Response

```json
success
```

### POST /report-quest  ~ ~  An adventurer report a quest

Request Body
```json
 {
    "adv_id": 1,
    "quest_id" : 1,
    "is_completed" : true

}
```

Response

```json
success
```