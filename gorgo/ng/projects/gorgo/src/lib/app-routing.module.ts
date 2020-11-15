import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';


import { ClassdiagramsTableComponent } from './classdiagrams-table/classdiagrams-table.component'
import { ClassdiagramAdderComponent } from './classdiagram-adder/classdiagram-adder.component'
import { ClassdiagramDetailComponent } from './classdiagram-detail/classdiagram-detail.component'
import { ClassdiagramPresentationComponent } from './classdiagram-presentation/classdiagram-presentation.component'

import { ClassshapesTableComponent } from './classshapes-table/classshapes-table.component'
import { ClassshapeAdderComponent } from './classshape-adder/classshape-adder.component'
import { ClassshapeDetailComponent } from './classshape-detail/classshape-detail.component'
import { ClassshapePresentationComponent } from './classshape-presentation/classshape-presentation.component'

import { FieldsTableComponent } from './fields-table/fields-table.component'
import { FieldAdderComponent } from './field-adder/field-adder.component'
import { FieldDetailComponent } from './field-detail/field-detail.component'
import { FieldPresentationComponent } from './field-presentation/field-presentation.component'

import { GorgoactionsTableComponent } from './gorgoactions-table/gorgoactions-table.component'
import { GorgoactionAdderComponent } from './gorgoaction-adder/gorgoaction-adder.component'
import { GorgoactionDetailComponent } from './gorgoaction-detail/gorgoaction-detail.component'
import { GorgoactionPresentationComponent } from './gorgoaction-presentation/gorgoaction-presentation.component'

import { LinksTableComponent } from './links-table/links-table.component'
import { LinkAdderComponent } from './link-adder/link-adder.component'
import { LinkDetailComponent } from './link-detail/link-detail.component'
import { LinkPresentationComponent } from './link-presentation/link-presentation.component'

import { PkgeltsTableComponent } from './pkgelts-table/pkgelts-table.component'
import { PkgeltAdderComponent } from './pkgelt-adder/pkgelt-adder.component'
import { PkgeltDetailComponent } from './pkgelt-detail/pkgelt-detail.component'
import { PkgeltPresentationComponent } from './pkgelt-presentation/pkgelt-presentation.component'

import { PositionsTableComponent } from './positions-table/positions-table.component'
import { PositionAdderComponent } from './position-adder/position-adder.component'
import { PositionDetailComponent } from './position-detail/position-detail.component'
import { PositionPresentationComponent } from './position-presentation/position-presentation.component'

import { StatesTableComponent } from './states-table/states-table.component'
import { StateAdderComponent } from './state-adder/state-adder.component'
import { StateDetailComponent } from './state-detail/state-detail.component'
import { StatePresentationComponent } from './state-presentation/state-presentation.component'

import { UmlscsTableComponent } from './umlscs-table/umlscs-table.component'
import { UmlscAdderComponent } from './umlsc-adder/umlsc-adder.component'
import { UmlscDetailComponent } from './umlsc-detail/umlsc-detail.component'
import { UmlscPresentationComponent } from './umlsc-presentation/umlsc-presentation.component'

import { VerticesTableComponent } from './vertices-table/vertices-table.component'
import { VerticeAdderComponent } from './vertice-adder/vertice-adder.component'
import { VerticeDetailComponent } from './vertice-detail/vertice-detail.component'
import { VerticePresentationComponent } from './vertice-presentation/vertice-presentation.component'

const routes: Routes = [
	{ path: 'classdiagrams', component: ClassdiagramsTableComponent, outlet: 'table' },
	{ path: 'classdiagram-adder', component: ClassdiagramAdderComponent, outlet: 'editor' },
	{ path: 'classdiagram-detail/:id', component: ClassdiagramDetailComponent, outlet: 'editor' },
	{ path: 'classdiagram-presentation/:id', component: ClassdiagramPresentationComponent, outlet: 'presentation' },
	{ path: 'classdiagram-presentation-special/:id', component: ClassdiagramPresentationComponent, outlet: 'classdiagrampres' },

	{ path: 'classshapes', component: ClassshapesTableComponent, outlet: 'table' },
	{ path: 'classshape-adder', component: ClassshapeAdderComponent, outlet: 'editor' },
	{ path: 'classshape-detail/:id', component: ClassshapeDetailComponent, outlet: 'editor' },
	{ path: 'classshape-presentation/:id', component: ClassshapePresentationComponent, outlet: 'presentation' },
	{ path: 'classshape-presentation-special/:id', component: ClassshapePresentationComponent, outlet: 'classshapepres' },

	{ path: 'fields', component: FieldsTableComponent, outlet: 'table' },
	{ path: 'field-adder', component: FieldAdderComponent, outlet: 'editor' },
	{ path: 'field-detail/:id', component: FieldDetailComponent, outlet: 'editor' },
	{ path: 'field-presentation/:id', component: FieldPresentationComponent, outlet: 'presentation' },
	{ path: 'field-presentation-special/:id', component: FieldPresentationComponent, outlet: 'fieldpres' },

	{ path: 'gorgoactions', component: GorgoactionsTableComponent, outlet: 'table' },
	{ path: 'gorgoaction-adder', component: GorgoactionAdderComponent, outlet: 'editor' },
	{ path: 'gorgoaction-detail/:id', component: GorgoactionDetailComponent, outlet: 'editor' },
	{ path: 'gorgoaction-presentation/:id', component: GorgoactionPresentationComponent, outlet: 'presentation' },
	{ path: 'gorgoaction-presentation-special/:id', component: GorgoactionPresentationComponent, outlet: 'gorgoactionpres' },

	{ path: 'links', component: LinksTableComponent, outlet: 'table' },
	{ path: 'link-adder', component: LinkAdderComponent, outlet: 'editor' },
	{ path: 'link-detail/:id', component: LinkDetailComponent, outlet: 'editor' },
	{ path: 'link-presentation/:id', component: LinkPresentationComponent, outlet: 'presentation' },
	{ path: 'link-presentation-special/:id', component: LinkPresentationComponent, outlet: 'linkpres' },

	{ path: 'pkgelts', component: PkgeltsTableComponent, outlet: 'table' },
	{ path: 'pkgelt-adder', component: PkgeltAdderComponent, outlet: 'editor' },
	{ path: 'pkgelt-detail/:id', component: PkgeltDetailComponent, outlet: 'editor' },
	{ path: 'pkgelt-presentation/:id', component: PkgeltPresentationComponent, outlet: 'presentation' },
	{ path: 'pkgelt-presentation-special/:id', component: PkgeltPresentationComponent, outlet: 'pkgeltpres' },

	{ path: 'positions', component: PositionsTableComponent, outlet: 'table' },
	{ path: 'position-adder', component: PositionAdderComponent, outlet: 'editor' },
	{ path: 'position-detail/:id', component: PositionDetailComponent, outlet: 'editor' },
	{ path: 'position-presentation/:id', component: PositionPresentationComponent, outlet: 'presentation' },
	{ path: 'position-presentation-special/:id', component: PositionPresentationComponent, outlet: 'positionpres' },

	{ path: 'states', component: StatesTableComponent, outlet: 'table' },
	{ path: 'state-adder', component: StateAdderComponent, outlet: 'editor' },
	{ path: 'state-detail/:id', component: StateDetailComponent, outlet: 'editor' },
	{ path: 'state-presentation/:id', component: StatePresentationComponent, outlet: 'presentation' },
	{ path: 'state-presentation-special/:id', component: StatePresentationComponent, outlet: 'statepres' },

	{ path: 'umlscs', component: UmlscsTableComponent, outlet: 'table' },
	{ path: 'umlsc-adder', component: UmlscAdderComponent, outlet: 'editor' },
	{ path: 'umlsc-detail/:id', component: UmlscDetailComponent, outlet: 'editor' },
	{ path: 'umlsc-presentation/:id', component: UmlscPresentationComponent, outlet: 'presentation' },
	{ path: 'umlsc-presentation-special/:id', component: UmlscPresentationComponent, outlet: 'umlscpres' },

	{ path: 'vertices', component: VerticesTableComponent, outlet: 'table' },
	{ path: 'vertice-adder', component: VerticeAdderComponent, outlet: 'editor' },
	{ path: 'vertice-detail/:id', component: VerticeDetailComponent, outlet: 'editor' },
	{ path: 'vertice-presentation/:id', component: VerticePresentationComponent, outlet: 'presentation' },
	{ path: 'vertice-presentation-special/:id', component: VerticePresentationComponent, outlet: 'verticepres' },

];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
