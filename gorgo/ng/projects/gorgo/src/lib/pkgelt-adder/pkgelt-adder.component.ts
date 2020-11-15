import { Component, OnInit } from '@angular/core';
import {FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { PkgeltDB } from '../pkgelt-db'
import { PkgeltService } from '../pkgelt.service'


@Component({
  selector: 'app-pkgelt-adder',
  templateUrl: './pkgelt-adder.component.html',
  styleUrls: ['./pkgelt-adder.component.css']
})
export class PkgeltAdderComponent implements OnInit {

	pkgelt = {} as PkgeltDB;






  constructor(
    private pkgeltService: PkgeltService, 
	  ) {
  }

  ngOnInit(): void {

  }


  add(): void {



    this.pkgeltService.postPkgelt( this.pkgelt )
    .subscribe(pkgelt => {
		this.pkgeltService.PkgeltServiceChanged.next("post")
		
		this.pkgelt = {} // reset fields
	    console.log("pkgelt added")
    });
  }
}
