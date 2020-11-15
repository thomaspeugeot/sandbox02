import { Component, OnInit } from '@angular/core';
import {FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { ClassdiagramDB } from '../classdiagram-db'
import { ClassdiagramService } from '../classdiagram.service'


import { Router, RouterState, ActivatedRoute } from '@angular/router';

@Component({
	selector: 'app-classdiagram-presentation',
	templateUrl: './classdiagram-presentation.component.html',
	styleUrls: ['./classdiagram-presentation.component.css']
})
export class ClassdiagramPresentationComponent implements OnInit {

	classdiagram: ClassdiagramDB;



	PkgeltsViaClassdiagramsFieldName = "Classdiagrams"; // Label used to generates the table of Pkgelt that points to Classdiagram via Classdiagrams
	PkgeltsViaClassdiagramsStructName = "Classdiagram"; // Label used to generates the table of Pkgelt that points to Classdiagram via Classdiagrams

	constructor(
		private classdiagramService: ClassdiagramService,

		private route: ActivatedRoute,
		private router: Router,
	) {
			this.router.routeReuseStrategy.shouldReuseRoute = function () {
				return false;
			};
	}

	ngOnInit(): void {
		this.getClassdiagram();


		// observable for changes in 
		this.classdiagramService.ClassdiagramServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getClassdiagram()
					
				}
			}
		)
	}

  getClassdiagram(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		this.classdiagramService.getClassdiagram(id)
		.subscribe( 
			classdiagram => 
			{ 
					this.classdiagram = classdiagram
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
	   			editor: ["classdiagram-detail", ID]
	 	}
   	}]);
 }

}
