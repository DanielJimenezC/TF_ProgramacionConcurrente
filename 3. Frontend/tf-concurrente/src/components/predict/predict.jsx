import React, { Component } from "react";
import Form from "react-bootstrap/Form";
import Button from "react-bootstrap/Button";
import { Row, Col, Grid, InputGroup } from "react-bootstrap";
import "./predict.css";
import { withRouter } from "react-router-dom";
import axios from "axios";

class Predict extends Component {
	state = {
		id: 1300,
		isAuthenticated: false,
		edad: "",
		peso: "",
		distrito: "3",
		tos: "0",
		fiebre: "0",
		dificultadRespira: "0",
		perdidaOlfato: "0",
		enfermo: "",
		responsePOST: "None",
	};

	handleChange = (event) => {
		const target = event.target;
		const value = target.value;
		const name = target.name;

		this.setState({
			[name]: value,
		});
	};

	submitForm = () => {
		axios
			.post("http://localhost:5000/api/users/prediction", this.state)
			.then((resonse) => {
				console.log(resonse);
				this.setState({ responsePOST: resonse.data });
				if (this.state.responsePOST == "true") {
					this.setState({ enfermo: "Enferma" });
				} else {
					this.setState({ enfermo: "No enferma" });
				}
			})
			.catch((error) => {
				console.log(error);
			});
	};

	render() {
		return (
			<React.Fragment>
				<Form className="styleForForm">
					<h5 className="text">Datos Personales</h5>
					<Form.Row className="row">
						<Col>
							<Form.Label>Edad</Form.Label>
							<InputGroup className="mb-2">
								<Form.Control
									type="text"
									name="edad"
									value={this.state.edad}
									onChange={this.handleChange}
								/>
							</InputGroup>
						</Col>
						<Col>
							<Form.Label>Peso</Form.Label>
							<InputGroup className="mb-2">
								<InputGroup.Prepend>
									<InputGroup.Text>Kg</InputGroup.Text>
								</InputGroup.Prepend>
								<Form.Control
									type="text"
									name="peso"
									value={this.state.peso}
									onChange={this.handleChange}
								/>
							</InputGroup>
						</Col>
					</Form.Row>
					<Form.Row className="row">
						<Col>
							<div class="form-group">
								<label for="exampleFormControlSelect1">Distrito</label>
								<select class="form-control" id="formSelect1">
									<option>Callao</option>
									<option>Ventanilla</option>
									<option>Ate</option>
									<option>Barranco</option>
									<option>Chorrillos</option>
									<option>Comas</option>
									<option>Jesus Maria</option>
									<option>La Molina</option>
									<option>La Victoria</option>
									<option>Lince</option>
									<option>Los Olivos</option>
									<option>Lurin</option>
									<option>Magdalena del Mar</option>
									<option>Miraflores</option>
									<option>Pueblo Libre</option>
									<option>Puente Piedra</option>
									<option>Rimac</option>
									<option>San Borja</option>
									<option>San Isidro</option>
									<option>San Juan de Lurigancho</option>
									<option>San Martin de Porres</option>
									<option>San Miguel</option>
									<option>Santiago de Surco</option>
									<option>Surquillo</option>
									<option>Villa El Salvador</option>
								</select>
							</div>
						</Col>
					</Form.Row>
					<h5 className="text">Sintomas</h5>
					<Form.Row className="row">
						<Col>
							<div class="form-group">
								<label for="exampleFormControlSelect1">¿Presentas tos?</label>
								<select class="form-control" id="formSelect2">
									<option>Verdad</option>
									<option>Falso</option>
								</select>
							</div>
						</Col>
					</Form.Row>
					<Form.Row className="row">
						<Col>
							<div class="form-group">
								<label for="exampleFormControlSelect1">
									¿Presentas fiebre?
								</label>
								<select class="form-control" id="formSelect2">
									<option>Verdad</option>
									<option>Falso</option>
								</select>
							</div>
						</Col>
					</Form.Row>
					<Form.Row className="row">
						<Col>
							<div class="form-group">
								<label for="exampleFormControlSelect1">
									¿Presentas dificultad de respirar?
								</label>
								<select class="form-control" id="formSelect2">
									<option>Verdad</option>
									<option>Falso</option>
								</select>
							</div>
						</Col>
					</Form.Row>
					<Form.Row className="row">
						<Col>
							<div class="form-group">
								<label for="exampleFormControlSelect1">
									¿Presentas perdida de gusto u olfato?
								</label>
								<select class="form-control" id="formSelect2">
									<option>Verdad</option>
									<option>Falso</option>
								</select>
							</div>
						</Col>
					</Form.Row>
				</Form>
				<div className="centerdiv">
					<Button
						variant="primary"
						className="button"
						onClick={this.submitForm}
					>
						Enviar
					</Button>
					<h6 className="textdiff">La persona esta:</h6>
					<span className="textResponse">{this.state.enfermo}</span>
				</div>
			</React.Fragment>
		);
	}
}

export default withRouter(Predict);
