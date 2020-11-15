import { Component, OnInit } from '@angular/core';

import { Router } from '@angular/router';

import { Observable, combineLatest, timer } from 'rxjs'

import { Paths } from './app-routing.module'
 
import * as laundromat from 'laundromat'
import * as gorgo from 'gorgo'

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {
  title = 'ng';

  diagramIDForWasherStates: number
  diagramIDForMachineStates: number

  constructor(
    private washersService: laundromat.WasherService,
    private machinesService: laundromat.MachineService,
    private umlscService: gorgo.UmlscService,

    private router: Router) {
  }

  // callbak function that is attached to the generic engine
  engineUpdatedCallbackFunction = (): void => {
    // console.log("engineUpdatedCallbackFunction called")

    this.machinesService.MachineServiceChanged.next("update")
    this.washersService.WasherServiceChanged.next("update")
    this.umlscService.UmlscServiceChanged.next("update")
  }

  ngOnInit(): void {

    combineLatest([
      this.umlscService.getUmlscs(),
    ]
    ).subscribe(
      ([
        umlscs,
      ]) => {
        umlscs.forEach(umlsc => {
          if (umlsc.Name.includes("Machine")) {
            this.diagramIDForMachineStates = umlsc.ID
          }
          if (umlsc.Name.includes("Washer")) {
            this.diagramIDForWasherStates = umlsc.ID
          }
        })

        // setup routers with machine & washer
        this.router.navigate([{
          outlets: {
            machinepres: [Paths.MACHINE_PRESENTATION_SPECIAL_PATH, 1],
            washerpres: [Paths.WASHER_PRESENTATION_SPECIAL_PATH, 1],

            machinestates: [ Paths.UMLSC_MACHINE_STATES_PATH, this.diagramIDForMachineStates],
            washerstates: [ Paths.UMLSC_WASHER_STATES_PATH, this.diagramIDForWasherStates],
          }

        }]);
      })
  }
}
