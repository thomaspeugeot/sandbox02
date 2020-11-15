import { Component, OnInit } from '@angular/core';
import {FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { MarshallerDB } from '../marshaller-db'
import { MarshallerService } from '../marshaller.service'


import { Router, RouterState, ActivatedRoute } from '@angular/router';

@Component({
	selector: 'app-marshaller-presentation',
	templateUrl: './marshaller-presentation.component.html',
	styleUrls: ['./marshaller-presentation.component.css']
})
export class MarshallerPresentationComponent implements OnInit {

	marshaller: MarshallerDB;




	constructor(
		private marshallerService: MarshallerService,

		private route: ActivatedRoute,
		private router: Router,
	) {
			this.router.routeReuseStrategy.shouldReuseRoute = function () {
				return false;
			};
	}

	ngOnInit(): void {
		this.getMarshaller();


		// observable for changes in 
		this.marshallerService.MarshallerServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getMarshaller()
					
				}
			}
		)
	}

  getMarshaller(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		this.marshallerService.getMarshaller(id)
		.subscribe( 
			marshaller => 
			{ 
					this.marshaller = marshaller
        	}
  	);
  }



	// set presentation outlet
	setPresentationRouterOutlet(structName :string, ID: number) {
		this.router.navigate([{
	  	outlets: {
			presentation: [structName + "-presentation", ID]
	  	}
		}]);
	}

	// set editor outlet
	setEditorRouterOutlet(ID: number) {
		this.router.navigate([{
	 		outlets: {
	   			editor: ["marshaller-detail", ID]
	 	}
   	}]);
 }

}
