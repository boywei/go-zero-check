package define

var (
	StmResult = `{
        "declaration": "",
        "automatons": [
          {
            "name": "SM1",
            "parameters": "c1, c2, c3, c4, c5 bool",
            "locations": [
              {
                "id": 1,
                "name": "A",
                "invariant": ""
              },
              {
                "id": 2,
                "name": "B",
                "invariant": ""
              },
              {
                "id": 3,
                "name": "C",
                "invariant": ""
              },
              {
                "id": 4,
                "name": "D",
                "invariant": ""
              },
              {
                "id": 5,
                "name": "E",
                "invariant": ""
              },
              {
                "id": 6,
                "name": "",
                "invariant": ""
              },
              {
                "id": 7,
                "name": "",
                "invariant": ""
              }
            ],
            "transitions": [
              {
                "id": 1,
                "source_id": 1,
                "target_id": 6,
                "select": "",
                "guard": "c1",
                "sync": "",
                "update": "S1()"
              },
              {
                "id": 2,
                "source_id": 6,
                "target_id": 7,
                "select": "",
                "guard": "!c2",
                "sync": "",
                "update": "S3()"
              },
              {
                "id": 3,
                "source_id": 6,
                "target_id": 2,
                "select": "",
                "guard": "c2",
                "sync": "",
                "update": "S2()"
              },
              {
                "id": 4,
                "source_id": 7,
                "target_id": 3,
                "select": "",
                "guard": "c4",
                "sync": "",
                "update": "S4()"
              },
              {
                "id": 5,
                "source_id": 7,
                "target_id": 4,
                "select": "",
                "guard": "!c4 && c5",
                "sync": "",
                "update": "S5()"
              },
              {
                "id": 6,
                "source_id": 7,
                "target_id": 5,
                "select": "",
                "guard": "!c4 && !c5",
                "sync": "",
                "update": "S6()"
              }
            ],
            "init": 1,
            "declaration": "func S1(){}\n\nfunc S2(){}\n\nfunc S3(){}\n\nfunc S4(){}\n\nfunc S5(){}\n\nfunc S6(){}"
          }
        ],
        "system_declaration": "system A;"
      }`
	CarResult = `{
  "declaration": "",
  "automatons": [
    {
      "name": "Car",
      "parameters": "",
      "init": 1,
      "declaration": "var v = 100\\n\\nfunc speed_up() {\\n    v = v + 10\\n}\\n\\nfunc speed_down() {\\n    v = v - 10\\n}\\n\\nfunc keep() {\\n    \\n}",
      "locations": [
        {
          "id": 3,
          "name": "Keep",
          "invariant": ""
        },
        {
          "id": 4,
          "name": "SpeedDown",
          "invariant": ""
        },
        {
          "id": 6,
          "name": "SpeedUp",
          "invariant": ""
        }
      ],
      "transitions": [
        {
          "id": 5,
          "source_id": 6,
          "target_id": 3,
          "select": "",
          "guard": "v <= 150 && v >= 50",
          "sync": "",
          "update": ""
        },
        {
          "id": 6,
          "source_id": 3,
          "target_id": 4,
          "select": "",
          "guard": "v >= 100",
          "sync": "",
          "update": ""
        },
        {
          "id": 11,
          "source_id": 6,
          "target_id": 4,
          "select": "",
          "guard": "",
          "sync": "",
          "update": "v >= 150"
        },
        {
          "id": 12,
          "source_id": 3,
          "target_id": 6,
          "select": "",
          "guard": "v <= 100",
          "sync": "",
          "update": ""
        },
        {
          "id": 13,
          "source_id": 6,
          "target_id": 6,
          "select": "",
          "guard": "",
          "sync": "",
          "update": "speed_up()"
        },
        {
          "id": 14,
          "source_id": 4,
          "target_id": 4,
          "select": "",
          "guard": "",
          "sync": "",
          "update": "speed_down()"
        },
        {
          "id": 15,
          "source_id": 3,
          "target_id": 3,
          "select": "",
          "guard": "",
          "sync": "",
          "update": "keep()"
        },
        {
          "id": 16,
          "source_id": 4,
          "target_id": 3,
          "select": "",
          "guard": "v >= 50 && v <= 150",
          "sync": "",
          "update": ""
        },
        {
          "id": 17,
          "source_id": 4,
          "target_id": 6,
          "select": "",
          "guard": "",
          "sync": "",
          "update": "v <= 100"
        }
      ]
    }
  ],
  "system_declaration": "// 在这里填写模型的声明.\\ncar := Car()"
}`
	NuclearResult = `{
  "declaration": "// 在这里编写全局变量、函数等.\\nvar (\\n    MIN_HEAT = 1000\\n    MAX_HEAT = 10000\\n    NORMAL_HEAT = 500\\n    COOL_HEAT = 400\\n)\\n\\nvar start, shutdown, emerg, cool, power_up, power_down bool",
  "automatons": [
    {
      "name": "Nuclear",
      "parameters": "",
      "init": 1,
      "declaration": "// 在这里编写局部变量、函数等.\\nvar Heat = 30\\n\\nfunc increase_heat() {\\n    Heat = Heat + 10\\n}\\n\\nfunc decrease_heat() {\\n    Heat = Heat - 10\\n}\\n\\nfunc decrease_heat_to_normal() {\\n    if Heat >= NORMAL_HEAT {\\n        Heat = Heat - 10\\n    }\\n}\\n\\nfunc increase_heat_to_min() {\\n    if Heat <= MIN_HEAT {\\n        Heat = Heat + 10\\n    }\\n}\\n\\nfunc keep_emergency() {\\n    \\n}\\n\\nfunc decrease_heat_to_cool() {\\n    if Heat >= COOL_HEAT {\\n        Heat = Heat - 10\\n    }\\n}\\n\\nfunc increase_heat_to_critical() {\\n    if Heat <= MIN_HEAT {\\n        Heat = Heat + 10\\n    }\\n}",
      "locations": [
        {
          "id": 1,
          "name": "ShutDown",
          "invariant": ""
        },
        {
          "id": 2,
          "name": "StartUp",
          "invariant": ""
        },
        {
          "id": 3,
          "name": "Cooling",
          "invariant": ""
        },
        {
          "id": 4,
          "name": "Emergency",
          "invariant": ""
        },
        {
          "id": 5,
          "name": "Critical",
          "invariant": ""
        },
        {
          "id": 6,
          "name": "PowerDown",
          "invariant": ""
        },
        {
          "id": 7,
          "name": "PowerUp",
          "invariant": ""
        }
      ],
      "transitions": [
        {
          "id": 1,
          "source_id": 1,
          "target_id": 2,
          "select": "",
          "guard": "",
          "sync": "shutdown?",
          "update": ""
        },
        {
          "id": 2,
          "source_id": 2,
          "target_id": 5,
          "select": "",
          "guard": "Heat >= MIN_HEAT",
          "sync": "",
          "update": "increase_heat()"
        },
        {
          "id": 3,
          "source_id": 5,
          "target_id": 7,
          "select": "",
          "guard": "",
          "sync": "power_up?",
          "update": "increase_heat()"
        },
        {
          "id": 4,
          "source_id": 7,
          "target_id": 6,
          "select": "",
          "guard": "Heat >= MAX_HEAT",
          "sync": "",
          "update": "decrease_heat()"
        },
        {
          "id": 5,
          "source_id": 6,
          "target_id": 3,
          "select": "",
          "guard": "",
          "sync": "cool?",
          "update": "decrease_heat()"
        },
        {
          "id": 6,
          "source_id": 3,
          "target_id": 4,
          "select": "",
          "guard": "",
          "sync": "emerg?",
          "update": ""
        },
        {
          "id": 7,
          "source_id": 2,
          "target_id": 4,
          "select": "",
          "guard": "",
          "sync": "emerg?",
          "update": ""
        },
        {
          "id": 8,
          "source_id": 1,
          "target_id": 4,
          "select": "",
          "guard": "",
          "sync": "emerg?",
          "update": ""
        },
        {
          "id": 9,
          "source_id": 3,
          "target_id": 1,
          "select": "",
          "guard": "",
          "sync": "shutdown?",
          "update": ""
        },
        {
          "id": 10,
          "source_id": 2,
          "target_id": 1,
          "select": "",
          "guard": "",
          "sync": "start?",
          "update": ""
        },
        {
          "id": 11,
          "source_id": 6,
          "target_id": 4,
          "select": "",
          "guard": "",
          "sync": "emerg?",
          "update": ""
        },
        {
          "id": 12,
          "source_id": 7,
          "target_id": 4,
          "select": "",
          "guard": "",
          "sync": "emerg?",
          "update": ""
        },
        {
          "id": 13,
          "source_id": 5,
          "target_id": 4,
          "select": "",
          "guard": "",
          "sync": "emerg?",
          "update": ""
        },
        {
          "id": 14,
          "source_id": 6,
          "target_id": 7,
          "select": "",
          "guard": "Heat <= MIN_HEAT",
          "sync": "",
          "update": "increase_heat()"
        },
        {
          "id": 15,
          "source_id": 7,
          "target_id": 3,
          "select": "",
          "guard": "",
          "sync": "cool?",
          "update": ""
        },
        {
          "id": 16,
          "source_id": 1,
          "target_id": 1,
          "select": "",
          "guard": "Heat <= NORMAL_HEAT",
          "sync": "",
          "update": "decrease_heat_to_normal()"
        },
        {
          "id": 17,
          "source_id": 2,
          "target_id": 2,
          "select": "",
          "guard": "Heat <= MIN_HEAT",
          "sync": "",
          "update": "increase_heat_to_min()"
        },
        {
          "id": 19,
          "source_id": 7,
          "target_id": 7,
          "select": "",
          "guard": "",
          "sync": "power_up?",
          "update": "increase_heat()"
        },
        {
          "id": 20,
          "source_id": 4,
          "target_id": 4,
          "select": "",
          "guard": "",
          "sync": "emergency?",
          "update": "keep_emergency()"
        },
        {
          "id": 21,
          "source_id": 3,
          "target_id": 3,
          "select": "",
          "guard": "Heat <= COOL_HEAT",
          "sync": "",
          "update": "decrease_heat_to_cool()"
        },
        {
          "id": 22,
          "source_id": 6,
          "target_id": 6,
          "select": "",
          "guard": "",
          "sync": "power_down?",
          "update": "decrease_heat()"
        },
        {
          "id": 23,
          "source_id": 5,
          "target_id": 5,
          "select": "",
          "guard": "Heat <= MAX_HEAT",
          "sync": "",
          "update": "increase_heat_to_critical()"
        },
        {
          "id": 24,
          "source_id": 4,
          "target_id": 2,
          "select": "",
          "guard": "",
          "sync": "start?",
          "update": ""
        }
      ]
    }
  ],
  "system_declaration": "// 在这里填写模型的声明.\\nnuclear := Nuclear()"
}`
)
