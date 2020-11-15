import { Component, OnInit } from '@angular/core';
import { Observable, combineLatest, timer } from 'rxjs';
import * as joint from 'jointjs';

import { ActivatedRoute, Router } from '@angular/router';

import * as gorgo from 'gorgo'
import { state } from '@angular/animations';

@Component({
  selector: 'app-umlsc-diagram',
  templateUrl: './umlsc-diagram.component.html',
  styleUrls: ['./umlsc-diagram.component.css']
})
export class UmlscDiagramComponent implements OnInit {

  namespace = joint.shapes;
  private paper: joint.dia.Paper;
  private graph = new joint.dia.Graph(
    {},
    { cellNamespace: this.namespace } // critical piece of code for save/restore diagrams with the jointjs stuff
  );

  public founded = false

  // the diagram of interest
  public umlscDB: gorgo.UmlscDB;

  // map states of the diagram
  mapStateDBIDStateDBs = new Map<number, gorgo.StateDB>()

  // map of States according to the joint.shapes.uml.State
  // it is used to save the diagram (which only know the ids)
  public MapJointjsIdsStates = new Map<string, gorgo.StateDB>();
  public MapStateDBIDJointjsStateID = new Map<number, string>()

  public MapNamesStates = new Map<string, gorgo.StateDB>()

  // current active state
  public activeState = ""

  constructor(
    private route: ActivatedRoute,
    private router: Router,
    private UmlscService: gorgo.UmlscService,
    private StateService: gorgo.StateService,
    private GorgoactionsServe: gorgo.GorgoactionService,
  ) {
    // https://stackoverflow.com/questions/54627478/angular-7-routing-to-same-component-but-different-param-not-working
    // this is for routerLink on same component when only queryParameter changes
    this.router.routeReuseStrategy.shouldReuseRoute = function () {
      return false;
    }
  }

  // if true the save button will appear
  public savebutton: boolean

  ngOnInit(): void {
    // wait for all fetch to combine
    const id = +this.route.snapshot.paramMap.get('id');
    if (this.route.snapshot.paramMap.has('savebutton')) {
      this.savebutton = (this.route.snapshot.paramMap.get('savebutton') == "true")
    }

    combineLatest([
      this.UmlscService.getUmlsc(id),
      this.StateService.getStates()
    ]).subscribe(
      ([umlscDB, stateDBs]) => {
        this.umlscDB = umlscDB

        // redeem state to state chart
        this.umlscDB.States = new Array<gorgo.StateDB>()
        stateDBs.forEach(stateDB => {
          if (stateDB.Umlsc_StatesDBID == this.umlscDB.ID) {
            this.mapStateDBIDStateDBs.set(stateDB.ID, stateDB)
            this.umlscDB.States.push(stateDB)
          }
        })
      })

    // draw diagram is only possible when the document element can be identified
    // with the current state chart (in case they are more than one)
    // Emits ascending numbers, one every second (1000ms), starting after 3 seconds
    timer(1000, 1000).subscribe(
      x => {
        var eltFromId = document.getElementById(this.umlscDB.Name)
        if (eltFromId != undefined && !this.founded) {
          this.founded = true
          this.drawClassdiagram()
        }
      }
    )

    // redraw the active state
    timer(2000, 1000).subscribe(
      x => {

        // redraw if active state changed
        this.UmlscService.getUmlsc(this.umlscDB.ID).subscribe(
          umlscDB => {
            if (this.activeState != umlscDB.Activestate) {
              this.activeState = umlscDB.Activestate
              this.updateDiagramWithColoredActiveState()
            }
          }
        )
      }
    )
  }

  updateDiagramWithColoredActiveState() {

    var cells = this.graph.getCells()
    // console.log(cells.length)
    var defaultColor = 'rgba(48, 208, 198, 0.1)'
    var redColor = 'rgba(248, 0, 0, 0.3)'

    cells.forEach(
      cell => {
        var cellId: any
        cellId = cell.id;
        if (this.MapJointjsIdsStates.get(cellId) != undefined) {

          // retrieve the shape.
          var stateDB = this.MapJointjsIdsStates.get(cellId)

          if (this.activeState == stateDB.Name) {
            cell.attr('.uml-state-body', { fill: redColor })
          } else {
            cell.attr('.uml-state-body', { fill: defaultColor })
          }
        }
      }
    )
  }

  drawClassdiagram(): void {
    const namespace = joint.shapes;
    this.paper = new joint.dia.Paper(
      {
        el: document.getElementById(this.umlscDB.Name),
        model: this.graph,
        width: 4000,
        // height: window.innerHeight,
        height: 4000,
        gridSize: 10,
        drawGrid: false,
        cellViewNamespace: namespace
      }
    )

    // redraw all states
    this.mapStateDBIDStateDBs.forEach(stateDB => {

      var color = 'rgba(48, 208, 198, 0.1)'

      if (stateDB.Name == this.activeState) {
        color = 'rgba(248, 0, 0, 0.3)'
      }

      var umlState = new joint.shapes.uml.State(
        {
          position: {
            x: stateDB.X,
            y: stateDB.Y
          },
          size: { width: 240, height: 40 },
          name: [stateDB.Name],
          attrs: {
            '.uml-state-body': {
              fill: color,
            },
          }
        })

      umlState.addTo(this.graph)

      // init coloring of state
      this.activeState = this.umlscDB.Activestate
      this.updateDiagramWithColoredActiveState()

      var id: any;
      id = umlState.id;
      var idstring: string
      idstring = id;
      this.MapJointjsIdsStates.set(idstring, stateDB)
      this.MapStateDBIDJointjsStateID.set(stateDB.ID, idstring)
    })
  }

  // later, we should do this with
  //
  // Listening for changes of the position to a single element
  // element1.on('change:position', function(element1, position) {
  //   alert('element1 moved to ' + position.x + ',' + position.y);
  // });
  saveClassdiagram(): void {
    console.log("save diagram")

    // parse shapes positions
    var cells = this.graph.getCells()
    console.log(cells.length)

    cells.forEach(
      cell => {
        // ugly hack because cell.id is considered a Dimension by the ts compiler
        // vive golang
        var cellId: any
        cellId = cell.id;
        if (this.MapJointjsIdsStates.get(cellId) != undefined) {

          // retrieve the shape.
          var stateDB = this.MapJointjsIdsStates.get(cellId)

          stateDB.X = cell.attributes.position.x
          stateDB.Y = cell.attributes.position.y

          // update position to DB
          this.StateService.updateState(stateDB).subscribe(
            position => {
              console.log("position updated")
            }
          )
        }
      }
    )

    // post SAVE Gorgoaction
    this.GorgoactionsServe.postGorgoaction(
      {
        Name: gorgo.Typeaction.MARSHALL_ALL_DIAGRAMS,
      } as gorgo.GorgoactionAPI
    ).subscribe(
      Gorgoaction => {
        console.log("Gorgoaction posted")
      }
    )
  }

}

