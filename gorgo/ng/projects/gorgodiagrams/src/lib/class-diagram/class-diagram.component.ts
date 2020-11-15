import { Component, OnInit } from '@angular/core';
import { Observable, combineLatest } from 'rxjs';
import * as joint from 'jointjs';

import { ActivatedRoute, Router } from '@angular/router';

import * as gorgo from 'gorgo'

@Component({
  selector: 'lib-class-diagram',
  templateUrl: './class-diagram.component.html',
  styleUrls: ['./class-diagram.component.css']
})
export class ClassDiagramComponent implements OnInit {

  namespace = joint.shapes;
  private paper: joint.dia.Paper;
  private graph = new joint.dia.Graph(
    {},
    { cellNamespace: this.namespace } // critical piece of code. 
  );

  // the diagram of interest
  public classdiagramDB: gorgo.ClassdiagramDB;

  public Classshapes = new Map<number, gorgo.ClassshapeDB>();
  public Links = new Map<number, gorgo.LinkDB>();
  public Fields = new Map<number, gorgo.FieldDB>();
  public Positions = new Map<number, gorgo.PositionDB>();
  public Vertices = new Map<number, gorgo.VerticeDB>();

  // map of Classhapes according to the joint.shapes.uml.Class
  // it is used to save the diagram (which only know the ids)
  public MapIdsClassshapes = new Map<string, gorgo.ClassshapeDB>();
  public MapClassshapenameClass = new Map<string, joint.shapes.uml.Class>();

  public MapIdsLinks = new Map<string, gorgo.LinkDB>();
  public MapLinksClass = new Map<string, joint.shapes.standard.Link>();

  constructor(
    private route: ActivatedRoute,
    private router: Router,
    private ClassshapeService: gorgo.ClassshapeService,
    private LinkService: gorgo.LinkService,
    private FieldService: gorgo.FieldService,
    private ClassdiagramService: gorgo.ClassdiagramService,
    private PositionService: gorgo.PositionService,
    private VerticeService: gorgo.VerticeService,
    private GorgoactionsServe: gorgo.GorgoactionService,
  ) {
    // https://stackoverflow.com/questions/54627478/angular-7-routing-to-same-component-but-different-param-not-working
    // this is for routerLink on same component when only queryParameter changes
    this.router.routeReuseStrategy.shouldReuseRoute = function () {
      return false;
    };
  }

  ngOnInit(): void {
    // ref: https://resources.jointjs.com/tutorial/hello-world
    // remember to add JointJS CSS in `angular.json#L33,99` (for `drawGrid` and others)
    // fetch data

    // wait for all fetch to combine
    const id = +this.route.snapshot.paramMap.get('id');

    combineLatest([
      this.ClassdiagramService.getClassdiagram(id),
      this.ClassshapeService.getClassshapes(),
      this.LinkService.getLinks(),
      this.FieldService.getFields(),
      this.PositionService.getPositions(),
      this.VerticeService.getVertices()
    ]).subscribe(
      ([
        classdiagram,
        classshapes,
        links,
        fields,
        positions,
        vertices,
      ]) => {

        this.classdiagramDB = classdiagram
        classshapes.forEach(classshape => this.Classshapes.set(classshape.ID, classshape))
        links.forEach(link => this.Links.set(link.ID, link))
        fields.forEach(field => this.Fields.set(field.ID, field))
        positions.forEach(position => this.Positions.set(position.ID, position))
        vertices.forEach(vertice => this.Vertices.set(vertice.ID, vertice))

        // redeem elements classshapes to classdiagram & position
        this.classdiagramDB.Classshapes = new Array<gorgo.ClassshapeDB>();
        this.Classshapes.forEach(
          classshape => {
            if (classshape.Classdiagram_ClassshapesDBID == this.classdiagramDB.ID) {
              this.classdiagramDB.Classshapes.push(classshape)
            }
            classshape.Position = this.Positions.get(classshape.PositionID)

            // redeem links to classshape
            classshape.Links = new Array<gorgo.LinkDB>();
            this.Links.forEach(
              link => {
                if (link.Classshape_LinksDBID == classshape.ID) {
                  classshape.Links.push(link)
                }
                link.Middlevertice = this.Vertices.get(link.MiddleverticeID)
              }
            )

            // redeem elements fields to classdiagram & position
            classshape.Fields = new Array<gorgo.FieldDB>();
            this.Fields.forEach(
              field => {
                if (field.Classshape_FieldsDBID == classshape.ID) {
                  classshape.Fields.push(field)
                }
              }
            )

          }
        )


        this.drawClassdiagram();
      }
    )
  }

  // to be completed
  drawClassdiagram(): void {
    console.log("draw diagram")

    const namespace = joint.shapes;
    this.paper = new joint.dia.Paper(
      {
        el: document.getElementById('jointjs-holder'),
        model: this.graph,
        width: 4000,
        // height: window.innerHeight,
        height: 4000,
        gridSize: 10,
        drawGrid: false,
        cellViewNamespace: namespace
      }
    );

    // draw class shapes
    this.classdiagramDB.Classshapes.forEach(classshape => {

      // fetch the fields, it must belong to the current diagram
      // and the type must match the classshape type
      var attributes = new Array<string>()
      classshape.Fields.forEach(
        field => {
          console.log(field.Fieldname + " " + field.Structname + " " + field.Fieldtypename)
          attributes.push(field.Fieldname + " : " + field.Fieldtypename)
        }
      )

      const umlClassShape = new joint.shapes.uml.Class(
        {
          position: {
            x: classshape.Position.X,
            y: classshape.Position.Y
          },
          size: { width: classshape.Width, height: classshape.Heigth },
          name: [classshape.Structname],
          attributes: attributes,
          methods: [],
          attrs: {
            '.uml-class-name-rect': {
              fill: '#ff8450',
              stroke: '#fff',
              'stroke-width': 0.5,
            },
            '.uml-class-name-text': {
              'font-family': 'Roboto'
          },
            '.uml-class-attrs-rect': {
              fill: '#fe976a',
              stroke: '#fff',
              height: 10,
              'stroke-width': 0.5,
              'font-family': 'Roboto'
            },
            '.uml-class-methods-rect': {
              fill: '#fe976a',
              stroke: '#fff',
              height: 0,
              'stroke-width': 0
            },
            '.uml-class-attrs-text': {
              'ref-y': 0,
              'y-alignment': 'top',
              'font-family': 'Roboto'
            }
          }
        }
      );

      // structRectangle.attributes = ['firstName: String']
      umlClassShape.addTo(this.graph);

      // horrible hack because the TS compiles assets that umlclasshape.id is not a string but a 
      // an attribute of type joint.dia.Dimension
      var id: any;
      id = umlClassShape.id;
      var idstring: string
      idstring = id;
      this.MapIdsClassshapes.set(idstring, classshape)
      // console.log("id " + umlClassShape.id + " idstring " + idstring)

      //
      this.MapClassshapenameClass.set(classshape.Structname, umlClassShape)
    })


    // draw links of the diagram shapes
    this.classdiagramDB.Classshapes.forEach(classshape => {

      classshape.Links.forEach(linkDB => {

        // does from & to shapes exists?
        var fromShape = this.MapClassshapenameClass.get(linkDB.Structname)
        var toShape = this.MapClassshapenameClass.get(linkDB.Fieldtypename)

        var strockWidth = 2
        let LinkEndlabel = linkDB.Fieldname
        let distance = 0.75

        let xFrom = fromShape.get('position').x
        let yFrom = fromShape.get('position').y
        let xTo = toShape.get('position').x
        let yTo = toShape.get('position').y
        var vertices = [{ x: (xFrom + yTo) / 2, y: (yFrom + yTo) / 2 }]

        if (linkDB.Middlevertice != undefined) {
          vertices = [{ x: linkDB.Middlevertice.X, y: linkDB.Middlevertice.Y }]
        }


        if (fromShape != undefined && toShape != undefined) {
          var link = new joint.shapes.standard.Link({
            source: fromShape,
            target: toShape,
            vertices: vertices,
            attrs: {
              line: {
                stroke: '#3c4260',
                strokeWidth: strockWidth,
                vertexMarker: {
                  'type': 'circle',
                  'r': 3,
                  'stroke-width': 2,
                  'fill': 'white'
                },
                targetMarker: { // no arrow at the end
                  'type': 'path',
                  'd': 'M 10 -5 0 0 10 5 z'
                },
              },
            },
            labels: [
              {
                attrs: { text: { text: LinkEndlabel } },
                position: {
                  offset: 15,
                  distance: distance
                }
              }
            ],
          })
          link.addTo(this.graph);

          // later, we need to save the diagram
          // 
          // algo is 
          // - for each cells of the diagram, 
          //      get its id & position
          //      find the original LinkDB and updates its position
          //
          // Because each cell has an unique id
          // we create a map of cell id to LinkDB in order 
          // horrible hack because the TS compiles assets that umlclasshape.id is not a string but a 
          // an attribute of type joint.dia.Dimension
          var id: any;
          id = link.id;
          var idstring: string
          idstring = id;
          this.MapIdsLinks.set(idstring, linkDB)
        }
      })
    })
  }

  // to be completed
  saveClassdiagram(): void {
    console.log("save diagram")

    // parse shapes positions
    var cells = this.graph.getCells()
    console.log(cells.length)

    cells.forEach(
      cell => {
        // ugly hack because cell.id is considered a Dimension by the ts compiler
        // vive golang
        var link: gorgo.ClassshapeDB;
        var cellId: any
        cellId = cell.id;
        if (this.MapIdsClassshapes.get(cellId) != undefined) {

          // retrieve the shape.
          link = this.MapIdsClassshapes.get(cellId)

          // fetch corresponding position and update
          var positionDB = this.Positions.get(link.PositionID)

          positionDB.X = cell.attributes.position.x
          positionDB.Y = cell.attributes.position.y

          // update position to DB
          this.PositionService.updatePosition(positionDB).subscribe(
            position => {
              console.log("position updated")
            }
          )
        }
        if (this.MapIdsLinks.has(cellId)) {

          // retrieve the shape.
          var linkDB = this.MapIdsLinks.get(cellId)

          // fetch corresponding position and update
          linkDB.Middlevertice.X = cell.attributes.vertices[0].x
          linkDB.Middlevertice.Y = cell.attributes.vertices[0].y

          // update position to DB
          var verticeDB = linkDB.Middlevertice
          this.VerticeService.updateVertice(verticeDB).subscribe(
            position => {
              console.log("vertice updated")
            }
          )
        }
      }
    )

    // post SAVE Gorgoaction
    this.GorgoactionsServe.postGorgoaction(
      {
        Name: "MARSHALL_ALL_DIAGRAMS",
      } as gorgo.GorgoactionAPI
    ).subscribe(
      Gorgoaction => {
        console.log("Gorgoaction posted")
      }
    )
  }
}

