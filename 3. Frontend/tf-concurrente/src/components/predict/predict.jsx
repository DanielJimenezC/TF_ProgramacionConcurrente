import React, { Component } from "react";
import Form from "react-bootstrap/Form";
import Button from "react-bootstrap/Button";
import { Row, Col, Grid, InputGroup } from "react-bootstrap";
import "./predict.css";
import { withRouter } from "react-router-dom";
import axios from "axios";
import Modal from "react-bootstrap/Modal";

var bgColors = { primary: "#1abc9c" };

class Predict extends Component {
	state = {
		id: 1300,
		isAuthenticated: false,
		edad: "",
		peso: "",
		distrito: "3",
		tos: "0",
		fiebre: "0",
		dificultadRespirar: "0",
		perdidaOlfato: "0",
		enfermo: "",
		chance: "",
		valueOfModal: false,
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
				this.setState({
					valueOfModal: true,
					chance: resonse.data.chance,
					valueOfModal: true,
				});
				if (resonse.data.enfermo) {
					this.setState({
						enfermo: "Enfermo",
					});
				} else {
					this.setState({
						enfermo: "Sano",
					});
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
								<select
									class="form-control"
									id="formSelect1"
									name="distrito"
									onClick={this.handleChange}
								>
									<option value="1">Callao</option>
									<option value="2">Ventanilla</option>
									<option value="3">Ate</option>
									<option value="4">Barranco</option>
									<option value="5">Chorrillos</option>
									<option value="6">Comas</option>
									<option value="7">Jesus Maria</option>
									<option value="8">La Molina</option>
									<option value="9">La Victoria</option>
									<option value="10">Lince</option>
									<option value="11">Los Olivos</option>
									<option value="12">Lurin</option>
									<option value="13">Magdalena del Mar</option>
									<option value="14">Miraflores</option>
									<option value="15">Pueblo Libre</option>
									<option value="16">Puente Piedra</option>
									<option value="17">Rimac</option>
									<option value="18">San Borja</option>
									<option value="19">San Isidro</option>
									<option value="20">San Juan de Lurigancho</option>
									<option value="21">San Martin de Porres</option>
									<option value="22">San Miguel</option>
									<option value="23">Santiago de Surco</option>
									<option value="24">Surquillo</option>
									<option value="25">Villa El Salvador</option>
								</select>
							</div>
						</Col>
					</Form.Row>
					<h5 className="text">Sintomas</h5>
					<Form.Row className="row">
						<Col>
							<div class="form-group">
								<label for="exampleFormControlSelect1">¿Presentas tos?</label>
								<select
									class="form-control"
									id="formSelect2"
									name="tos"
									onClick={this.handleChange}
								>
									<option value="1">Verdad</option>
									<option value="0">Falso</option>
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
								<select
									class="form-control"
									id="formSelect2"
									name="fiebre"
									onClick={this.handleChange}
								>
									<option value="1">Verdad</option>
									<option value="0">Falso</option>
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
								<select
									class="form-control"
									id="formSelect2"
									name="dificultadRespirar"
									onClick={this.handleChange}
								>
									<option value="1">Verdad</option>
									<option value="0">Falso</option>
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
								<select
									class="form-control"
									id="formSelect2"
									name="perdidaOlfato"
									onClick={this.handleChange}
								>
									<option value="1">Verdad</option>
									<option value="0">Falso</option>
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
				</div>
				<Modal
					size="lg"
					aria-labelledby="contained-modal-title-vcenter"
					centered
					show={this.state.valueOfModal}
					onHide={() => this.setState({ valueOfModal: false })}
				>
					<Modal.Header closeButton>
						<Modal.Title id="contained-modal-title-vcenter">
							Resultado
						</Modal.Title>
					</Modal.Header>
					<div className="text-center">
						<h6>Hay un {this.state.chance}% de chance que usted este</h6>
						<span style={{ color: bgColors.primary }}>
							{this.state.enfermo}
						</span>
					</div>
					<Modal.Body className="text-center">
						<p>
							Esto no remplaza un examen real, y usted debería buscar ayuda
							profesional.
						</p>
						<a href="/faq">Ver como prevenir contagio</a>
					</Modal.Body>
				</Modal>
			</React.Fragment>
		);
	}
}

export default withRouter(Predict);
