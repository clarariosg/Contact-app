package main

import (
	"slices"
	"strings"
)

type Contact struct {
	ID int
 	First string
	Last string
	Email string
	Phone string
} 

func search_contact_by_id(id int) *Contact {
	for i := range contactList {
		if contactList[i].ID == id {
			return &contactList[i]
		}
	}
	return nil
}

func delete_contact_by_id(id int) bool {
	for i := range contactList {
		if contactList[i].ID == id {
			contactList = slices.Delete(contactList, i, i+1)
			return true
		}
	}
	return false
}

func validate_email(id int, email string) string {
	if email == "" {
		return "Email vacío, escriba algo por favor"
	}

	if !strings.Contains(email, "@") || !strings.Contains(email, ".") {
		return "Formato no válido de email"
	}

	for _, c := range contactList {
		if c.Email == email && c.ID != id {
			return "Email en uso. Tiene que ser único"
		}
	}

	return ""
}

var contactList = []Contact{
	{ID: 1, First: "Clara", Last: "Rios", Phone: "644431064", Email: "clara@example.com"},
	{ID: 2, First: "Nacho", Last: "Rios", Phone: "622612006", Email: "nacho@example.com"},
	{ID: 3, First: "Lucas", Last: "Martin", Phone: "611234567", Email: "lucas.martin@example.com"},
	{ID: 4, First: "Sofia", Last: "Lopez", Phone: "612345678", Email: "sofia.lopez@example.com"},
	{ID: 5, First: "Daniel", Last: "Garcia", Phone: "613456789", Email: "daniel.garcia@example.com"},
	{ID: 6, First: "Marta", Last: "Sanchez", Phone: "614567890", Email: "marta.sanchez@example.com"},
	{ID: 7, First: "Hugo", Last: "Perez", Phone: "615678901", Email: "hugo.perez@example.com"},
	{ID: 8, First: "Laura", Last: "Gomez", Phone: "616789012", Email: "laura.gomez@example.com"},
	{ID: 9, First: "Pablo", Last: "Fernandez", Phone: "617890123", Email: "pablo.fernandez@example.com"},
	{ID: 10, First: "Elena", Last: "Ruiz", Phone: "618901234", Email: "elena.ruiz@example.com"},
	{ID: 11, First: "Mario", Last: "Diaz", Phone: "619012345", Email: "mario.diaz@example.com"},
	{ID: 12, First: "Lucia", Last: "Moreno", Phone: "620123456", Email: "lucia.moreno@example.com"},
	{ID: 13, First: "Adrian", Last: "Munoz", Phone: "621234567", Email: "adrian.munoz@example.com"},
	{ID: 14, First: "Paula", Last: "Alonso", Phone: "623456789", Email: "paula.alonso@example.com"},
	{ID: 15, First: "Diego", Last: "Gutierrez", Phone: "624567890", Email: "diego.gutierrez@example.com"},
	{ID: 16, First: "Irene", Last: "Navarro", Phone: "625678901", Email: "irene.navarro@example.com"},
	{ID: 17, First: "Javier", Last: "Torres", Phone: "626789012", Email: "javier.torres@example.com"},
	{ID: 18, First: "Alba", Last: "Dominguez", Phone: "627890123", Email: "alba.dominguez@example.com"},
	{ID: 19, First: "Sergio", Last: "Vazquez", Phone: "628901234", Email: "sergio.vazquez@example.com"},
	{ID: 20, First: "Nuria", Last: "Ramos", Phone: "629012345", Email: "nuria.ramos@example.com"},
	{ID: 21, First: "Alejandro", Last: "Gil", Phone: "630123456", Email: "alejandro.gil@example.com"},
	{ID: 22, First: "Carmen", Last: "Serrano", Phone: "631234567", Email: "carmen.serrano@example.com"},
	{ID: 23, First: "Raul", Last: "Blanco", Phone: "632345678", Email: "raul.blanco@example.com"},
	{ID: 24, First: "Andrea", Last: "Molina", Phone: "633456789", Email: "andrea.molina@example.com"},
	{ID: 25, First: "Victor", Last: "Suarez", Phone: "634567890", Email: "victor.suarez@example.com"},
	{ID: 26, First: "Sara", Last: "Ortega", Phone: "635678901", Email: "sara.ortega@example.com"},
	{ID: 27, First: "Alvaro", Last: "Delgado", Phone: "636789012", Email: "alvaro.delgado@example.com"},
	{ID: 28, First: "Beatriz", Last: "Castro", Phone: "637890123", Email: "beatriz.castro@example.com"},
	{ID: 29, First: "Miguel", Last: "Ortiz", Phone: "638901234", Email: "miguel.ortiz@example.com"},
	{ID: 30, First: "Cristina", Last: "Rubio", Phone: "639012345", Email: "cristina.rubio@example.com"},
	{ID: 31, First: "Ivan", Last: "Marin", Phone: "640123456", Email: "ivan.marin@example.com"},
	{ID: 32, First: "Teresa", Last: "Sanz", Phone: "641234567", Email: "teresa.sanz@example.com"},
	{ID: 33, First: "Oscar", Last: "Iglesias", Phone: "642345678", Email: "oscar.iglesias@example.com"},
	{ID: 34, First: "Ana", Last: "Medina", Phone: "643456789", Email: "ana.medina@example.com"},
	{ID: 35, First: "Ruben", Last: "Cortes", Phone: "645678901", Email: "ruben.cortes@example.com"},
	{ID: 36, First: "Silvia", Last: "Garrido", Phone: "646789012", Email: "silvia.garrido@example.com"},
	{ID: 37, First: "Marcos", Last: "Santos", Phone: "647890123", Email: "marcos.santos@example.com"},
	{ID: 38, First: "Patricia", Last: "Herrera", Phone: "648901234", Email: "patricia.herrera@example.com"},
	{ID: 39, First: "David", Last: "Cabrera", Phone: "649012345", Email: "david.cabrera@example.com"},
	{ID: 40, First: "Alicia", Last: "Vega", Phone: "650123456", Email: "alicia.vega@example.com"},
	{ID: 41, First: "Guillermo", Last: "Campos", Phone: "651234567", Email: "guillermo.campos@example.com"},
	{ID: 42, First: "Eva", Last: "Reyes", Phone: "652345678", Email: "eva.reyes@example.com"},
	{ID: 43, First: "Jorge", Last: "Pascual", Phone: "653456789", Email: "jorge.pascual@example.com"},
	{ID: 44, First: "Monica", Last: "Vicente", Phone: "654567890", Email: "monica.vicente@example.com"},
	{ID: 45, First: "Fernando", Last: "Mendez", Phone: "655678901", Email: "fernando.mendez@example.com"},
	{ID: 46, First: "Rocio", Last: "Calvo", Phone: "656789012", Email: "rocio.calvo@example.com"},
	{ID: 47, First: "Alberto", Last: "Nieto", Phone: "657890123", Email: "alberto.nieto@example.com"},
	{ID: 48, First: "Lorena", Last: "Duran", Phone: "658901234", Email: "lorena.duran@example.com"},
	{ID: 49, First: "Manuel", Last: "Prieto", Phone: "659012345", Email: "manuel.prieto@example.com"},
	{ID: 50, First: "Noelia", Last: "Vidal", Phone: "660123456", Email: "noelia.vidal@example.com"},
	{ID: 51, First: "Antonio", Last: "Leon", Phone: "661234567", Email: "antonio.leon@example.com"},
	{ID: 52, First: "Isabel", Last: "Moya", Phone: "662345678", Email: "isabel.moya@example.com"},
	{ID: 53, First: "Samuel", Last: "Parra", Phone: "663456789", Email: "samuel.parra@example.com"},
	{ID: 54, First: "Natalia", Last: "Bravo", Phone: "664567890", Email: "natalia.bravo@example.com"},
	{ID: 55, First: "Gonzalo", Last: "Roca", Phone: "665678901", Email: "gonzalo.roca@example.com"},
	{ID: 56, First: "Julia", Last: "Acosta", Phone: "666789012", Email: "julia.acosta@example.com"},
	{ID: 57, First: "Eduardo", Last: "Soler", Phone: "667890123", Email: "eduardo.soler@example.com"},
	{ID: 58, First: "Claudia", Last: "Pastor", Phone: "668901234", Email: "claudia.pastor@example.com"},
	{ID: 59, First: "Roberto", Last: "Esteban", Phone: "669012345", Email: "roberto.esteban@example.com"},
	{ID: 60, First: "Ines", Last: "Gallego", Phone: "670123456", Email: "ines.gallego@example.com"},
	{ID: 61, First: "Carlos", Last: "Roman", Phone: "671234567", Email: "carlos.roman@example.com"},
	{ID: 62, First: "Maria", Last: "Valero", Phone: "672345678", Email: "maria.valero@example.com"},
	{ID: 63, First: "Andres", Last: "Carrasco", Phone: "673456789", Email: "andres.carrasco@example.com"},
	{ID: 64, First: "Belen", Last: "Pardo", Phone: "674567890", Email: "belen.pardo@example.com"},
	{ID: 65, First: "Enrique", Last: "Merino", Phone: "675678901", Email: "enrique.merino@example.com"},
	{ID: 66, First: "Aitana", Last: "Lorenzo", Phone: "676789012", Email: "aitana.lorenzo@example.com"},
	{ID: 67, First: "Felix", Last: "Caballero", Phone: "677890123", Email: "felix.caballero@example.com"},
	{ID: 68, First: "Miriam", Last: "Escobar", Phone: "678901234", Email: "miriam.escobar@example.com"},
	{ID: 69, First: "Nicolas", Last: "Fuentes", Phone: "679012345", Email: "nicolas.fuentes@example.com"},
	{ID: 70, First: "Esther", Last: "Cano", Phone: "680123456", Email: "esther.cano@example.com"},
	{ID: 71, First: "Emilio", Last: "Bernal", Phone: "681234567", Email: "emilio.bernal@example.com"},
	{ID: 72, First: "Raquel", Last: "Segura", Phone: "682345678", Email: "raquel.segura@example.com"},
	{ID: 73, First: "Jaime", Last: "Arias", Phone: "683456789", Email: "jaime.arias@example.com"},
	{ID: 74, First: "Lidia", Last: "Crespo", Phone: "684567890", Email: "lidia.crespo@example.com"},
	{ID: 75, First: "Alex", Last: "Santana", Phone: "685678901", Email: "alex.santana@example.com"},
	{ID: 76, First: "Gloria", Last: "Miranda", Phone: "686789012", Email: "gloria.miranda@example.com"},
	{ID: 77, First: "Rafael", Last: "Peña", Phone: "687890123", Email: "rafael.pena@example.com"},
	{ID: 78, First: "Veronica", Last: "Redondo", Phone: "688901234", Email: "veronica.redondo@example.com"},
	{ID: 79, First: "Tomas", Last: "Izquierdo", Phone: "689012345", Email: "tomas.izquierdo@example.com"},
	{ID: 80, First: "Diana", Last: "Cuesta", Phone: "690123456", Email: "diana.cuesta@example.com"},
	{ID: 81, First: "Bruno", Last: "Benitez", Phone: "691234567", Email: "bruno.benitez@example.com"},
	{ID: 82, First: "Olga", Last: "Pineda", Phone: "692345678", Email: "olga.pineda@example.com"},
	{ID: 83, First: "Mateo", Last: "Montes", Phone: "693456789", Email: "mateo.montes@example.com"},
	{ID: 84, First: "Pilar", Last: "Rosales", Phone: "694567890", Email: "pilar.rosales@example.com"},
	{ID: 85, First: "Gabriel", Last: "Ferrer", Phone: "695678901", Email: "gabriel.ferrer@example.com"},
	{ID: 86, First: "Marina", Last: "Salas", Phone: "696789012", Email: "marina.salas@example.com"},
	{ID: 87, First: "Luis", Last: "Luna", Phone: "697890123", Email: "luis.luna@example.com"},
	{ID: 88, First: "Ainhoa", Last: "Saez", Phone: "698901234", Email: "ainhoa.saez@example.com"},
	{ID: 89, First: "Ignacio", Last: "Blazquez", Phone: "699012345", Email: "ignacio.blazquez@example.com"},
	{ID: 90, First: "Mireia", Last: "Navas", Phone: "600123456", Email: "mireia.navas@example.com"},
	{ID: 91, First: "Cesar", Last: "Roldan", Phone: "601234567", Email: "cesar.roldan@example.com"},
	{ID: 92, First: "Amalia", Last: "Costa", Phone: "602345678", Email: "amalia.costa@example.com"},
	{ID: 93, First: "Borja", Last: "Franco", Phone: "603456789", Email: "borja.franco@example.com"},
	{ID: 94, First: "Carla", Last: "Mora", Phone: "604567890", Email: "carla.mora@example.com"},
	{ID: 95, First: "Ricardo", Last: "Aguilar", Phone: "605678901", Email: "ricardo.aguilar@example.com"},
	{ID: 96, First: "Ariadna", Last: "Lozano", Phone: "606789012", Email: "ariadna.lozano@example.com"},
	{ID: 97, First: "Joaquin", Last: "Carmona", Phone: "607890123", Email: "joaquin.carmona@example.com"},
	{ID: 98, First: "Emma", Last: "Casado", Phone: "608901234", Email: "emma.casado@example.com"},
	{ID: 99, First: "Rodrigo", Last: "Mateos", Phone: "609012345", Email: "rodrigo.mateos@example.com"},
	{ID: 100, First: "Leire", Last: "Valle", Phone: "610123456", Email: "leire.valle@example.com"},
}