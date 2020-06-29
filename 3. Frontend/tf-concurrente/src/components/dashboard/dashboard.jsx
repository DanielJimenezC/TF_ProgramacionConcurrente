import React, { Component } from "react";
import Form from "react-bootstrap/Form";
import Button from "react-bootstrap/Button";
import { withRouter } from "react-router-dom";
import Card from "react-bootstrap/Card";
import CardDeck from "react-bootstrap/CardDeck";
import "./dashboard.css";
import CardColumns from "react-bootstrap/CardColumns";
class Dashboard extends Component {
	render() {
		return (
			<React.Fragment>
				<CardColumns
					style={{
						padding: "25px",
					}}
				>
					<Card>
						<Card.Img
							variant="top"
							src="https://www.paho.org/sites/default/files/styles/flexslider_full/public/hero/2020-03/covid-19-1190x574-2-full.jpg?h=fdc0eb87&itok=CerQCpzv"
						/>
						<Card.Body>
							<Card.Title>Examen Virtual de COVID-19</Card.Title>
							<Card.Text>
								This is a longer card with supporting text below as a natural
								lead-in to additional content. This content is a little bit
								longer.
							</Card.Text>
							<div className="text-center">
								<Button variant="primary">Dar Examen</Button>
							</div>
						</Card.Body>
					</Card>
					<Card>
						<Card.Img
							variant="top"
							src="https://espacio.fundaciontelefonica.com/wp-content/uploads/2019/08/block_chain_1100x525-1-1100x550.jpg"
						/>
						<Card.Body>
							<Card.Title>Registrar Datos</Card.Title>
							<Card.Text>
								This card has supporting text below as a natural lead-in to
								additional content.{" "}
							</Card.Text>
							<div className="text-center">
								<Button variant="primary">Dar Examen</Button>
							</div>
						</Card.Body>
					</Card>
					<Card className="text-center">
						<Card.Body>
							<Card.Title>Card title</Card.Title>
							<Card.Text>
								This card has supporting text below as a natural lead-in to
								additional content.{" "}
							</Card.Text>
							<div className="text-center">
								<Button variant="primary">Dar Examen</Button>
							</div>
						</Card.Body>

						<Card.Footer>
							<small className="text-muted">Last updated 3 mins ago</small>
						</Card.Footer>
					</Card>
					<Card>
						<Card.Img
							variant="top"
							src="https://elmontonero.pe/upload/uploads_images/columna_alan_6.jpg"
						/>
						<Card.Body>
							<Card.Title>Card title</Card.Title>
							<Card.Text>
								This card has supporting text below as a natural lead-in to
								additional content.{" "}
							</Card.Text>
						</Card.Body>
					</Card>
				</CardColumns>
			</React.Fragment>
		);
	}
}

export default withRouter(Dashboard);
