# tcc-golang

Projeto para matricula de uma faculdade

# Funcionalidades

No projeto é possivel
* Cadastrar um aluno
* Visualizar as informações de um aluno cadastrada
* Visualizar as informações de matricula um aluno
* Cadastrar uma materia
* Realizar a matricula

# Endpoints

> POST /v1/aluno

Request Body:
```
{
	"nome": "Guilherme",
	"documento" : "121271465"	
}
```
Response Body:
```
{
	"id": 5,
	"nome": "Guilherme",
	"documento": "121271465"
}
```

> GET /v1/aluno/:alunoID

Response Body:
```
{
	"id": 2,
	"nome": "Guilherme",
	"documento": "121271"
}
```


> GET /v1/aluno/:alunoID/matricula

Response Body:
```
{
	"matriculaDTO": [
		{
			"id": 3,
			"aluno_id": 2,
			"materia_id": 2,
			"descricao_semestre": "SM/6",
			"finalizado": false
		},
		{
			"id": 4,
			"aluno_id": 2,
			"materia_id": 3,
			"descricao_semestre": "SM/6",
			"finalizado": false
		}
	],
	"id": 2,
	"nome": "Guilherme",
	"documento": "121271"
}

```

>POST /v1/materia
Request Body:
```
{
	"nome": "Calculo 8",
	"capacidade": 35,
	"materia_dependencia": 2
}
```

Response Body:
```
{
	"id": 8,
	"nome": "Calculo 8",
	"capacidade": 35,
	"materia_dependencia": 2
}
```

>POST /v1/matricula
Request Body:
```
{
	"aluno_id": 2,
	"materia_id": 3,
	"descricao_semestre": "SM/6",
	"finalizado": false
}
```

Response Body:
```
{
	"id": 4,
	"aluno_id": 2,
	"materia_id": 3,
	"descricao_semestre": "SM/6",
	"finalizado": false
}
```

