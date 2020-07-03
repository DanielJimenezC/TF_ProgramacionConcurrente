import React, { Component } from "react";
import axios from "axios";
import { AgGridReact } from "ag-grid-react";
import "ag-grid-community/dist/styles/ag-grid.css";
import "ag-grid-community/dist/styles/ag-theme-alpine.css";
import Container from "react-bootstrap/Container";
import { Row, Col, Grid, InputGroup } from "react-bootstrap";
import "./groups.css";

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
			},
			{
				headerName: "Peso",
				field: "peso",
			},
			{
				headerName: "Distrito",
				field: "distrito",
			},
		],
	};

	componentDidMount() {
		axios
			.get("http://localhost:5000/api/groups")
			.then((resonse) => {
				console.log(resonse.data.clusters.Clusters[0]);
				this.setState(
					{
						clusters: resonse.data.clusters.Clusters,
						group1: resonse.data.clusters.Clusters[0],
						group2: resonse.data.clusters.Clusters[1],
						group3: resonse.data.clusters.Clusters[2],
						group4: resonse.data.clusters.Clusters[3],
					},
					() => {
						this.autoSizeAll(true);
					}
				);
			})
			.catch((error) => {
				console.log(error);
			});
	}

	onGridReady = (params) => {
		this.gridApi = params.api;
		this.gridColumnApi = params.columnApi;

		const httpRequest = new XMLHttpRequest();
		const updateData = (data) => {
			this.setState({ rowData: data });
		};

		httpRequest.open(
			"GET",
			"https://raw.githubusercontent.com/ag-grid/ag-grid/master/grid-packages/ag-grid-docs/src/olympicWinnersSmall.json"
		);
		httpRequest.send();
		httpRequest.onreadystatechange = () => {
			if (httpRequest.readyState === 4 && httpRequest.status === 200) {
				updateData(JSON.parse(httpRequest.responseText));
			}
		};
	};

	sizeToFit = () => {
		this.gridApi.sizeColumnsToFit();
	};

	autoSizeAll = (skipHeader) => {
		var allColumnIds = [];
		this.gridColumnApi.getAllColumns().forEach(function (column) {
			allColumnIds.push(column.colId);
		});
		this.gridColumnApi.autoSizeColumns(allColumnIds, skipHeader);
	};

	render() {
		return (
			<React.Fragment>
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
							}}
						>
							<h5>Grupo 1</h5>
							<div
								className="ag-theme-alpine"
								style={{
									height: "350px",
								}}
							>
								<AgGridReact
									modules={this.state.modules}
									columnDefs={this.state.columnDefs}
									rowData={this.state.group1.Data}
									onGridReady={this.onGridReady}
								></AgGridReact>
							</div>
						</Col>
						<Col>
							<h5>Grupo 2</h5>
							<div
								className="ag-theme-alpine"
								style={{
									height: "350px",
									width: "600",
								}}
							>
								<AgGridReact
									modules={this.state.modules}
									columnDefs={this.state.columnDefs}
									rowData={this.state.group2.Data}
									onGridReady={this.onGridReady}
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
									height: "350px",
									width: "600",
								}}
							>
								<AgGridReact
									modules={this.state.modules}
									columnDefs={this.state.columnDefs}
									rowData={this.state.group3.Data}
									onGridReady={this.onGridReady}
								></AgGridReact>
							</div>
						</Col>

						<Col>
							<h5>Grupo 4</h5>{" "}
							<div
								className="ag-theme-alpine"
								style={{
									height: "350px",
									width: "600",
								}}
							>
								<AgGridReact
									modules={this.state.modules}
									columnDefs={this.state.columnDefs}
									rowData={this.state.group4.Data}
									onGridReady={this.onGridReady}
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
