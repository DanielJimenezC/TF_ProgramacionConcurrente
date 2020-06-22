import React, { Component } from "react";
import Form from "react-bootstrap/Form";
import Button from "react-bootstrap/Button";
import { Row, Col, Grid, InputGroup } from "react-bootstrap";
import "./predict.css";

class Predict extends Component {
	state = {
		id: 1300,
		peso: "",
		isAuthenticated: false,
	};

	handleChange = (event) => {
		const target = event.target;
		const value = target.value;
		const name = target.name;

		this.setState({
			[name]: value,
		});
	};

	render() {
		return (
			<React.Fragment>
				<Form className="styleForForm">
					<h5 className="text">Datos Personales</h5>
					<Form.Row className="row">
						<Col>
							<Form.Label>Fecha de Nacimiento</Form.Label>
							<InputGroup className="mb-2">
								<Form.Control
									type="text"
									name="peso"
									value={this.state.peso}
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
									<option>1</option>
									<option>2</option>
								</select>
							</div>
						</Col>
					</Form.Row>
					<h5 className="text">Síntomas</h5>
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
				</div>
			</React.Fragment>
		);
	}
}

export default Predict;
