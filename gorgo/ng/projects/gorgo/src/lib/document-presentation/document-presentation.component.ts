import { Component, OnInit } from '@angular/core';
import {FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { DocumentDB } from '../document-db'
import { DocumentService } from '../document.service'


import { Router, RouterState, ActivatedRoute } from '@angular/router';

@Component({
	selector: 'app-document-presentation',
	templateUrl: './document-presentation.component.html',
	styleUrls: ['./document-presentation.component.css']
})
export class DocumentPresentationComponent implements OnInit {

	document: DocumentDB;



	PkgeltsViaDocumentsFieldName = "Documents"; // Label used to generates the table of Pkgelt that points to Document via Documents
	PkgeltsViaDocumentsStructName = "Document"; // Label used to generates the table of Pkgelt that points to Document via Documents

	constructor(
		private documentService: DocumentService,

		private route: ActivatedRoute,
		private router: Router,
	) {
			this.router.routeReuseStrategy.shouldReuseRoute = function () {
				return false;
			};
	}

	ngOnInit(): void {
		this.getDocument();


		// observable for changes in 
		this.documentService.DocumentServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getDocument()
					
				}
			}
		)
	}

  getDocument(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		this.documentService.getDocument(id)
		.subscribe( 
			document => 
			{ 
					this.document = document
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
	   			editor: ["document-detail", ID]
	 	}
   	}]);
 }

}
