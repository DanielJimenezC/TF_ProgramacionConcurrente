import React, { Component } from "react";
import Form from "react-bootstrap/Form";
import Button from "react-bootstrap/Button";
import { withRouter } from "react-router-dom";
import Card from "react-bootstrap/Card";
import CardDeck from "react-bootstrap/CardDeck";
import "./dashboard.css";
import CardColumns from "react-bootstrap/CardColumns";
import Overlay from "react-bootstrap/Overlay";
import OverlayTrigger from "react-bootstrap/OverlayTrigger";
import Tooltip from "react-bootstrap/Tooltip";

function renderTooltip(props) {
	return (
		<Tooltip id="button-tooltip" {...props}>
			Simple tooltip
		</Tooltip>
	);
}

class Dashboard extends Component {
	renderTooltip(props, type) {
		switch (type) {
			case "registrar":
				return (
					<Tooltip id="button-tooltip" {...props}>
						Esta consistencia de datos se logra mediante blockchain.
					</Tooltip>
				);
				break;
			case "examen":
				return (
					<Tooltip id="button-tooltip" {...props}>
						Esto no remplaza un examen real, solo es para ver si tienes una alta
						probabilidad de estar infectado.
					</Tooltip>
				);
				break;
		}
		return null;
	}
	render() {
		return (
			<React.Fragment>
				<CardColumns
					style={{
						padding: "25px",
					}}
				>
					{" "}
					<a
						style={{ cursor: "pointer" }}
						onClick={() => {
							this.props.history.push("/register");
						}}
					>
						<OverlayTrigger
							placement="right"
							delay={{ show: 150, hide: 100 }}
							overlay={this.renderTooltip(this.props, "registrar")}
						>
							<Card>
								<div class="inner">
									<Card.Img
										variant="top"
										src="https://espacio.fundaciontelefonica.com/wp-content/uploads/2019/08/block_chain_1100x525-1-1100x550.jpg"
									/>
								</div>
								<Card.Body>
									<Card.Title>Registrar Datos</Card.Title>
									<Card.Text>
										Registra de forma confiable los datos de personas que han
										sido infectadas por COVID-19.
									</Card.Text>
								</Card.Body>
							</Card>
						</OverlayTrigger>
					</a>
					<a
						style={{ cursor: "pointer" }}
						onClick={() => {
							this.props.history.push("/predict");
						}}
					>
						<OverlayTrigger
							placement="right"
							delay={{ show: 150, hide: 100 }}
							overlay={this.renderTooltip(this.props, "examen")}
						>
							<Card>
								<div class="inner">
									<Card.Img
										variant="top"
										src="https://www.paho.org/sites/default/files/styles/flexslider_full/public/hero/2020-03/covid-19-1190x574-2-full.jpg?h=fdc0eb87&itok=CerQCpzv"
									/>
								</div>
								<Card.Body>
									<Card.Title>Examen Virtual de Covid-19</Card.Title>
									<Card.Text>
										De un examen virtual de Covid-19 para ver si hay una
										probabilidad de que estés infectado.
									</Card.Text>
								</Card.Body>
							</Card>
						</OverlayTrigger>
					</a>
					<a
						style={{ cursor: "pointer" }}
						onClick={() => {
							this.props.history.push("/faq");
						}}
					>
						<Card>
							<Card.Body>
								<Card.Title>Preguntas Frecuentes</Card.Title>
								<Card.Text>
									Averigua la información relevante sobre la prevención y
									tratamiento del nuevo corona virus, dicho por la Organización
									Mundial de la Salud.
								</Card.Text>
							</Card.Body>
							<Card.Footer>
								<small className="text-muted">
									Actualizado hace 30 minutos
								</small>
							</Card.Footer>
						</Card>
					</a>
					<a
						style={{ cursor: "pointer" }}
						onClick={() => {
							this.props.history.push("/groups");
						}}
					>
						<Card>
							<div class="inner">
								<Card.Img
									variant="top"
									src="https://elmontonero.pe/upload/uploads_images/columna_alan_6.jpg"
								/>
							</div>
							<Card.Body>
								<Card.Title>Grupos de riesgo</Card.Title>
								<Card.Text>
									Averigua cuales son los grupos de riesgo ante el Covid-19.
								</Card.Text>
							</Card.Body>
						</Card>
					</a>
				</CardColumns>
			</React.Fragment>
		);
	}
}

export default withRouter(Dashboard);
