import { Component, OnInit } from '@angular/core';

import * as gorgo from 'gorgo'

@Component({
  selector: 'app-action-buttons',
  templateUrl: './action-buttons.component.html',
  styleUrls: ['./action-buttons.component.css']
})
export class ActionButtonsComponent implements OnInit {

  constructor(
    private GorgoactionService: gorgo.GorgoactionService
  ) { }

  ngOnInit(): void {
  }

  reload() {
    this.GorgoactionService.postGorgoaction( { Name: gorgo.Typeaction.UNMARSHALL_ALL_DIAGRAMS, } as gorgo.GorgoactionDB
    ).subscribe( action => { console.log("action posted") } )
  }
}
