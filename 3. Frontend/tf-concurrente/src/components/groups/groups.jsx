import React, { Component } from "react";
import axios from "axios";
import { AgGridReact } from "ag-grid-react";
import "ag-grid-community/dist/styles/ag-grid.css";
import "ag-grid-community/dist/styles/ag-theme-alpine.css";
import Container from "react-bootstrap/Container";
import { Row, Col, Grid, InputGroup } from "react-bootstrap";
import "./groups.css";
import Alert from "react-bootstrap/Alert";

class Groups extends Component {
	state = {
		isAuthenticated: false,
		userName: "",
		clusters: [],
		group1: [],
		group2: [],
		group3: [],
		group4: [],
		columnDefs: [
			{
				headerName: "Edad",
				field: "edad",
				width: 80,
			},
			{
				headerName: "Peso",
				field: "peso",
				width: 80,
			},
			{
				headerName: "Distrito",
				field: "distrito",
				width: 150,
			},
			{
				headerName: "Dif. al Respirar",
				field: "dificultadRespirar",

				width: 140,
			},
			{
				headerName: "Prd. de Olfato",
				field: "perdidaOlfato",

				width: 140,
			},
		],
	};

	componentDidMount() {
		axios
			.get("http://localhost:5000/api/groups")
			.then((resonse) => {
				console.log(resonse.data.clusters.Clusters[0]);
				this.setState({
					clusters: resonse.data.clusters.Clusters,
					group1: resonse.data.clusters.Clusters[0],
					group2: resonse.data.clusters.Clusters[1],
					group3: resonse.data.clusters.Clusters[2],
					group4: resonse.data.clusters.Clusters[3],
				});
			})
			.catch((error) => {
				console.log(error);
			});
	}

	render() {
		return (
			<React.Fragment>
				<Alert variant={"info"}>
					Clustering de grupos utilizando K-means, con un K constante de 4!
				</Alert>
				<Container>
					<Row
						style={{
							display: "flex",
							justifyContent: "space-between",
							padding: "15px",
						}}
					>
						<Col
							style={{
								paddingRight: "50px",
								width: "500px",
							}}
						>
							<h5>Grupo 1</h5>
							<div
								className="ag-theme-alpine"
								style={{
									height: "300px",
								}}
							>
								<AgGridReact
									columnDefs={this.state.columnDefs}
									rowData={this.state.group1.Data}
								></AgGridReact>
							</div>
						</Col>
						<Col>
							<h5>Grupo 2</h5>
							<div
								className="ag-theme-alpine"
								style={{
									height: "300px",
									width: "500px",
								}}
							>
								<AgGridReact
									columnDefs={this.state.columnDefs}
									rowData={this.state.group2.Data}
								></AgGridReact>
							</div>
						</Col>
					</Row>
					<Row
						style={{
							display: "flex",
							justifyContent: "space-between",
							padding: "15px",
						}}
					>
						<Col
							style={{
								paddingRight: "50px",
							}}
						>
							<h5>Grupo 3</h5>
							<div
								className="ag-theme-alpine"
								style={{
									height: "300px",
									width: "500px",
								}}
							>
								<AgGridReact
									columnDefs={this.state.columnDefs}
									rowData={this.state.group3.Data}
								></AgGridReact>
							</div>
						</Col>

						<Col>
							<h5>Grupo 4</h5>{" "}
							<div
								className="ag-theme-alpine"
								style={{
									height: "300px",
									width: "500px",
								}}
							>
								<AgGridReact
									columnDefs={this.state.columnDefs}
									rowData={this.state.group4.Data}
								></AgGridReact>
							</div>
						</Col>
					</Row>
				</Container>
			</React.Fragment>
		);
	}
}

export default Groups;
