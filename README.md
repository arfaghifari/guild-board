# guild-board

## IDEAS
Do u need a hand to help but dont know how to ask? Now, There is a GUILD BOARD. U can request a quest to guild, and our adventurer will go to help u. Dont forget to bring the commission after the quest have completed

## Build and Running
Install module
```
Go mod tidy
```
Build : 
```
Make build
```
Run :
```
Make run
```
Run test :
```
go test -cover ./...
```


## List API
### GET /quest-status  ~ ~ Get All Quest
Query : "status" = 0|1 

Body : {}

```json
{
    "header": {
        "error_code": "",
        "status_code": 200
    },
    "data": [
        {
            "quest_id": 3,
            "name": "mengusir ular dari rumah",
            "description": "keluar ular dari kamar mandi",
            "minimum_rank": 13,
            "reward_number": 700000
        },
        {
            "quest_id": 4,
            "name": "menjaga anak",
            "description": "menjaga anak 6 tahun selama sehari",
            "minimum_rank": 12,
            "reward_number": 500000
        }
    ]
}
```

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
{
    "header": {
        "error_code": "",
        "status_code": 200
    },
    "data": {
        "quest_id": 6,
        "name": "menjaga anak",
        "description": "menjaga anak 6 tahun selama sehari",
        "minimum_rank": 12,
        "reward_number": 500000,
        "status": 0
    }
}
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
{
    "header": {
        "error_code": "",
        "status_code": 200
    },
    "data": {
        "success": true
    }
}
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
{
    "header": {
        "error_code": "",
        "status_code": 200
    },
    "data": {
        "success": true
    }
}
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
{
    "header": {
        "error_code": "",
        "status_code": 200
    },
    "data": {
        "success": true
    }
}
```

### POST /adventurer  ~ ~ Make an adventurer

Request Body
```json
{
    "name" : "naufal",
    "rank" : 11
}
```

Response

```json
{
    "header": {
        "error_code": "",
        "status_code": 200
    },
    "data": {
        "id": 5,
        "name": "naufal",
        "rank": 12,
        "completed_quest": 1
    }
}
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
{
    "header": {
        "error_code": "",
        "status_code": 200
    },
    "data": {
        "success": true
    }
}
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
{
    "header": {
        "error_code": "",
        "status_code": 200
    },
    "data": {
        "success": true
    }
}
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
{
    "header": {
        "error_code": "",
        "status_code": 200
    },
    "data": {
        "success": true
    }
}
```

### GET /adventurer  ~ ~ Get adventurer
Query : "adv_id" > 0

Body : {}

```json
{
    "header": {
        "error_code": "",
        "status_code": 200
    },
    "data": {
        "id": 5,
        "name": "naufal",
        "rank": 12,
        "completed_quest": 1
    }
}
```

### GET /quest-active-adv  ~ ~ Get active quest adventurer
Query : "adv_id" > 0 

Body : {}

```json
{
    "header": {
        "error_code": "",
        "status_code": 200
    },
    "data": [
        {
            "quest_id": 3,
            "name": "mengusir ular dari rumah",
            "description": "keluar ular dari kamar mandi",
            "minimum_rank": 13,
            "reward_number": 700000,
            "status" : 1
        },
        {
            "quest_id": 4,
            "name": "menjaga anak",
            "description": "menjaga anak 6 tahun selama sehari",
            "minimum_rank": 12,
            "reward_number": 500000,
            "status" : 1
        }
    ]
}
```