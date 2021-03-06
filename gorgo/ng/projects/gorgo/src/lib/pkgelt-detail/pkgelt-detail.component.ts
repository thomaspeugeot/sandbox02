// generated by genNgDetail.go
import { Component, OnInit } from '@angular/core';
import {FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { PkgeltDB } from '../pkgelt-db'
import { PkgeltService } from '../pkgelt.service'



import { Router, RouterState, ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-pkgelt-detail',
  templateUrl: './pkgelt-detail.component.html',
  styleUrls: ['./pkgelt-detail.component.css']
})
export class PkgeltDetailComponent implements OnInit {

	// the PkgeltDB of interest
	pkgelt: PkgeltDB;

	constructor(
		private pkgeltService: PkgeltService,

		private route: ActivatedRoute,
		private router: Router,
	) {
		// https://stackoverflow.com/questions/54627478/angular-7-routing-to-same-component-but-different-param-not-working
		// this is for routerLink on same component when only queryParameter changes
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
		return false;
		};
  }

  ngOnInit(): void {
	this.getPkgelt();

  }

  getPkgelt(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		this.pkgeltService.getPkgelt(id)
		.subscribe( 
			pkgelt => 
			{ 
					this.pkgelt = pkgelt

        }
  	);
  }


  save(): void {
	const id = +this.route.snapshot.paramMap.get('id');




	this.pkgeltService.updatePkgelt( this.pkgelt )
    .subscribe(pkgelt => {
		this.pkgeltService.PkgeltServiceChanged.next("update")

    	console.log("pkgelt saved")
    });
  }
}
