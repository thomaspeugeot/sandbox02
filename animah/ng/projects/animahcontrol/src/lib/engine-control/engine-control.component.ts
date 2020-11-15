import { Component, OnInit, Input } from '@angular/core';

import * as animah from 'animah'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { Observable, combineLatest, timer } from 'rxjs'


@Component({
  selector: 'lib-engine-control',
  templateUrl: './engine-control.component.html',
  styleUrls: ['./engine-control.component.css']
})
export class EngineControlComponent implements OnInit {

  public engine: animah.EngineDB
  public engineID: number

  lastEvent: string;
  lastEventAgent: string;

  nextEventAgent: string;
  nextEventName: string;
  nextEventTime: string;

  engineEventNumber: number;

  // animation fof the simulation
  speed = 36;
  clientState = "PAUSED";

  currTime: number;
  obsTimer: Observable<number> = timer(1000, 1000);

  diagramIDForSamocStates: number
  diagramIDForNatoStates: number

  constructor(
    private engineService: animah.EngineService,
    private actionService: animah.ActionService,

    private router: Router) {
  }

  // engineUpdated is the call function
  @Input() engineUpdatedCallbackFunction : () => void;

  ngOnInit(): void {

    // get the current engine
    this.engineService.getEngines().subscribe(
      engines => {
        engines.forEach(
          engine => {
            this.engine = engine
            this.engineID = this.engine.ID
          }
        )
      }
    )


    this.obsTimer.subscribe(
      currTime => {
        this.currTime = currTime

        if (this.engine != undefined) {
          this.engineService.getEngine(this.engineID).subscribe(
            engine => {
              this.engine = engine

              // this is the callback function from the generic engien to the specific engine 
              if (this.engineUpdatedCallbackFunction != undefined) {
                this.engineUpdatedCallbackFunction()
              }
            }
          )
        }
      }
    )
  }

  fireEventTillStateChange(): void {
    this.actionService.postAction({ Name: animah.ActionType.FIRE_EVENT_TILL_STATES_CHANGE }).subscribe(
      action => {
        console.log("action " + action.Name)
      }
    )
  }

  fireEvent(): void {
    this.actionService.postAction({ Name: animah.ActionType.FIRE_NEXT_EVENT }).subscribe(
      action => {
        console.log("action " + action.Name)
      }
    )
  }

  reset(): void {
    this.actionService.postAction({ Name: animah.ActionType.RESET }).subscribe(
      action => {
        console.log("action " + action.Name)
      }
    )
  }

  play(): void {
    this.actionService.postAction({ Name: animah.ActionType.PLAY }).subscribe(
      action => {
        console.log("action " + action.Name)
      }
    )
  }

  pause(): void {
    this.actionService.postAction({ Name: animah.ActionType.PAUSE }).subscribe(
      action => {
        console.log("action " + action.Name)
      }
    )
  }

  increaseSpeed100percent() : void {
    this.actionService.postAction({ Name: animah.ActionType.INCREASE_SPEED_100_PERCENTS }).subscribe(
      action => {
        console.log("action " + action.Name)
      }
    )
  }

  decreaseSpeed50percent() : void {
    this.actionService.postAction({ Name: animah.ActionType.DECREASE_SPEED_50_PERCENTS }).subscribe(
      action => {
        console.log("action " + action.Name)
      }
    )    
  }
}
