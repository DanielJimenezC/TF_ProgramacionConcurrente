import React, { Component } from "react";
import Accordion from "react-bootstrap/Accordion";
import Card from "react-bootstrap/Card";
import Button from "react-bootstrap/Button";
import "./faq.css";
import ListGroup from "react-bootstrap/ListGroup";
class Faq extends Component {
	state = {
		isAuthenticated: false,
		userName: "",
	};

	render() {
		return (
			<React.Fragment>
				<div class="styleForForm">
					<h5 style={{ padding: "30px" }} className="text-center">
						Preguntas Frecuentes brindadas por la Organización Mundial de la
						Salud
					</h5>
					<Accordion>
						<Card>
							<Card.Header>
								<Accordion.Toggle as={Button} variant="Secondary" eventKey="0">
									¿Qué es un coronavirus?
								</Accordion.Toggle>
							</Card.Header>
							<Accordion.Collapse eventKey="0">
								<Card.Body>
									Los coronavirus son una amplia familia de virus que se
									encuentran tanto ‎en animales como en humanos. Algunos
									infectan al ser humano y se ‎sabe que pueden causar diversas
									afecciones, desde el resfriado común ‎hasta enfermedades más
									graves como el síndrome respiratorio de ‎Oriente Medio (MERS)
									y el síndrome respiratorio agudo severo (SRAS).‎
								</Card.Body>
							</Accordion.Collapse>
						</Card>
						<Card>
							<Card.Header>
								<Accordion.Toggle as={Button} variant="Secondary" eventKey="1">
									¿El nuevo coronavirus se propaga por medio de aerosoles?
								</Accordion.Toggle>
							</Card.Header>
							<Accordion.Collapse eventKey="1">
								<Card.Body>
									Cuando las personas quieren saber cómo protegerse contra
									enfermedades respiratorias, suelen surgir preguntas acerca de
									los aerosoles. Cuando alguien estornuda o tose, puede esparcir
									gotículas grandes, pero las gotículas no permanecen
									suspendidas en el aire durante mucho tiempo, sino que caen.
									Algunos procedimientos sanitarios como la intubación pueden
									esparcir pequeñas gotículas en el aire. Las gotículas de mayor
									tamaño caen rápidamente. Las más pequeñas caen más lentamente.
									Conocemos la existencia de la contaminación ambiental por
									MERS-CoV, ya que se ha encontrado ARN del virus en sistemas de
									filtración de aire (pero no el virus vivo). Sin embargo, en lo
									que respecta al nuevo coronavirus, todavía tenemos que
									analizar los datos y comprender cómo se ha evaluado la
									transmisión.
								</Card.Body>
							</Accordion.Collapse>
						</Card>
						<Card>
							<Card.Header>
								<Accordion.Toggle as={Button} variant="Secondary" eventKey="2">
									¿Cuánto dura el periodo de incubación?‎
								</Accordion.Toggle>
							</Card.Header>
							<Accordion.Collapse eventKey="2">
								<Card.Body>
									El periodo de incubación es el intervalo de tiempo que
									transcurre entre ‎la infección y la aparición de los síntomas
									clínicos de la enfermedad. Las ‎estimaciones actuales apuntan
									a que el periodo de incubación varía ‎entre 1 y 12,5 días, con
									una media estimada de 5-6 días. Estas estimaciones se irán
									ajustando a medida que se disponga de más ‎datos. Sobre la
									base de la información disponible sobre otras enfermedades
									provocadas por coronavirus, entre ellas el MERS y el SARS, se
									estima que el periodo de ‎incubación del 2019-nCoV podría ser
									hasta de 14 días. ‎La OMS recomienda que el seguimiento de
									contactos de casos confirmados sea de 14 días.
								</Card.Body>
							</Accordion.Collapse>
						</Card>
						<Card>
							<Card.Header>
								<Accordion.Toggle as={Button} variant="Secondary" eventKey="3">
									¿Qué puedo hacer para protegerme? ‎
								</Accordion.Toggle>
							</Card.Header>
							<Accordion.Collapse eventKey="3">
								<Card.Body>
									<ListGroup>
										<ListGroup.Item>
											1. Lávese las manos frecuentemente
										</ListGroup.Item>
										<ListGroup.Item>
											2. Adopte medidas de higiene respiratoria
										</ListGroup.Item>
										<ListGroup.Item>
											3. Mantenga el distanciamiento social
										</ListGroup.Item>
										<ListGroup.Item>
											4. Evite tocarse los ojos, la nariz y la boca
										</ListGroup.Item>
										<ListGroup.Item>
											5. Si tiene fiebre, tos y dificultad para respirar,
											solicite atención médica a tiempo
										</ListGroup.Item>
									</ListGroup>
								</Card.Body>
							</Accordion.Collapse>
						</Card>
						<Card>
							<Card.Header>
								<Accordion.Toggle as={Button} variant="Secondary" eventKey="4">
									¿Debo ponerme una mascarilla para protegerme?‎
								</Accordion.Toggle>
							</Card.Header>
							<Accordion.Collapse eventKey="4">
								<Card.Body>
									Usar una mascarilla médica puede ayudar a limitar la
									propagación de ‎algunas enfermedades respiratorias. Sin
									embargo, el uso de una ‎mascarilla no garantiza por sí solo
									que no se contraigan infecciones y ‎debe combinarse con otras
									medidas de prevención, en particular la ‎higiene respiratoria
									y de las manos y la evitación del contacto cercano – ‎debe
									guardar por lo menos 1 metro (3 pies) de distancia con las
									demás ‎personas. ‎ La OMS aconseja el uso racional de las
									mascarillas médicas para evitar el ‎derroche innecesario de
									recursos valiosos y su posible uso indebido ‎‎(véanse los
									consejos sobre la utilización de mascarillas). Esto significa
									‎que solo debe utilizar una mascarilla si presenta síntomas
									respiratorios ‎‎(tos o estornudos), si sospecha que tiene
									infección por el 2019-nCoV ‎con síntomas leves o si está
									cuidando de alguien de quien se sospeche ‎que está infectado
									por el 2019-nCoV. Se ha de sospechar una infección ‎por
									2019-nCoV si la persona en cuestión ha viajado a una zona de
									China ‎en la que se haya notificado la presencia del 2019-nCoV
									o si ha tenido ‎contacto cercano con alguien que haya viajado
									desde China y tenga ‎síntomas respiratorios.‎
								</Card.Body>
							</Accordion.Collapse>
						</Card>
						<Card>
							<Card.Header>
								<Accordion.Toggle as={Button} variant="Secondary" eventKey="5">
									‎¿Cuán peligroso es?‎
								</Accordion.Toggle>
							</Card.Header>
							<Accordion.Collapse eventKey="5">
								<Card.Body>
									Tal como ocurre con otras enfermedades respiratorias, la
									infección por ‎el 2019-nCoV puede causar síntomas leves como
									rinorrea, dolor de ‎garganta, tos y fiebre. En algunas
									personas puede ser más grave y ‎causar neumonía o dificultades
									respiratorias. En raras ocasiones, la ‎enfermedad puede ser
									mortal. Las personas de edad avanzada y las ‎personas con
									afecciones comórbidas (como diabetes o cardiopatías) ‎parecen
									correr un mayor riesgo de caer gravemente enfermas por el
									‎virus.‎
								</Card.Body>
							</Accordion.Collapse>
						</Card>
					</Accordion>
				</div>
			</React.Fragment>
		);
	}
}

export default Faq;
