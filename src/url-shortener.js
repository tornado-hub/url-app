import React,{Component} from "react";
import axios from "axios";
import {Cards, Header, Form, Input, Icon} from "semantic-ui-react";

let endpoint = "https://localhost:9000";

class UrlShortener extends  Component{
	constructor(props){
		super(props);
		this.state={
			URL: "",
			shortened:"",
		}
	}
	ComponentDidMount(){
		this.getTask();
	}
	render(){
		return(
				<div >
					<div className="row">
						<Header className="header" as="h2" color="yellow">
							To DO list
						</Header>
					</div>
				</div>
		);
	}
}

export default UrlShortener;