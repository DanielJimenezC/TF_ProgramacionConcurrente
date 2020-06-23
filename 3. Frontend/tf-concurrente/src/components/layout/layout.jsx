import React, { Component } from "react";
import { Navbar, Nav } from "react-bootstrap";

class Layout extends Component {
	state = {
		isAuthenticated: false,
		userName: "",
		isOpen: false,
	};

	logout = () => {
		localStorage.clear();
	};

	toggleOpen = () => this.setState({ isOpen: !this.state.isOpen });

	render() {
		const menuClass = `dropdown-menu${this.state.isOpen ? " show" : ""}`;
		return (
			<React.Fragment>
				<Navbar bg="light" expand="lg" bg="dark" variant="dark">
					<Navbar.Brand>Covid-19 Tool</Navbar.Brand>
					<Navbar.Toggle aria-controls="basic-navbar-nav" />
					<Navbar.Collapse className="justify-content-end">
						<Nav className="mr-auto">
							{Boolean(localStorage.getItem("auth")) && (
								<Nav.Link href="predict">Examen de Covid-19</Nav.Link>
							)}
							{Boolean(localStorage.getItem("auth")) && (
								<Nav.Link href="faq">Preguntas Frecuentes</Nav.Link>
							)}
						</Nav>
						<ul class="navbar-nav">
							{Boolean(localStorage.getItem("auth")) && (
								<div className="dropdown" onClick={this.toggleOpen}>
									<button
										className="btn btn-secondary dropdown-toggle"
										type="button"
										id="dropdownMenuButton"
										data-toggle="dropdown"
										aria-haspopup="true"
									>
										Signed in as: {localStorage.getItem("username")}
									</button>
									<div
										className={menuClass}
										aria-labelledby="dropdownMenuButton"
									>
										<a className="dropdown-item" href="/" onClick={this.logout}>
											Logout
										</a>
									</div>
								</div>
							)}
						</ul>
					</Navbar.Collapse>
				</Navbar>
			</React.Fragment>
		);
	}
}

export default Layout;
