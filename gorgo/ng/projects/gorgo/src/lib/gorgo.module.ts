import { NgModule } from '@angular/core';

import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { Routes, RouterModule } from '@angular/router';

// for angular material
import { MatSliderModule } from '@angular/material/slider';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatSelectModule } from '@angular/material/select'
import { MatDatepickerModule } from '@angular/material/datepicker'
import { MatTableModule } from '@angular/material/table'
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatToolbarModule } from '@angular/material/toolbar'
import { MatListModule } from '@angular/material/list'
import { MatExpansionModule } from '@angular/material/expansion'; 

import { AngularSplitModule, SplitComponent } from 'angular-split';

import { AppRoutingModule } from './app-routing.module';

import { SplitterComponent} from './splitter/splitter.component'
import { SidebarComponent } from './sidebar/sidebar.component';


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


@NgModule({
  declarations: [


	ClassdiagramsTableComponent,
	ClassdiagramAdderComponent,
	ClassdiagramDetailComponent,
	ClassdiagramPresentationComponent,

	ClassshapesTableComponent,
	ClassshapeAdderComponent,
	ClassshapeDetailComponent,
	ClassshapePresentationComponent,

	FieldsTableComponent,
	FieldAdderComponent,
	FieldDetailComponent,
	FieldPresentationComponent,

	GorgoactionsTableComponent,
	GorgoactionAdderComponent,
	GorgoactionDetailComponent,
	GorgoactionPresentationComponent,

	LinksTableComponent,
	LinkAdderComponent,
	LinkDetailComponent,
	LinkPresentationComponent,

	PkgeltsTableComponent,
	PkgeltAdderComponent,
	PkgeltDetailComponent,
	PkgeltPresentationComponent,

	PositionsTableComponent,
	PositionAdderComponent,
	PositionDetailComponent,
	PositionPresentationComponent,

	StatesTableComponent,
	StateAdderComponent,
	StateDetailComponent,
	StatePresentationComponent,

	UmlscsTableComponent,
	UmlscAdderComponent,
	UmlscDetailComponent,
	UmlscPresentationComponent,

	VerticesTableComponent,
	VerticeAdderComponent,
	VerticeDetailComponent,
	VerticePresentationComponent,


	SplitterComponent,
	SidebarComponent
  ],
  imports: [
	FormsModule,
	ReactiveFormsModule,
    CommonModule,
    RouterModule,

    AppRoutingModule,

    MatSliderModule,
    MatSelectModule,
    MatFormFieldModule,
    MatInputModule,
    MatDatepickerModule,
    MatTableModule,
    MatCheckboxModule,
    MatButtonModule,
    MatIconModule,
	MatToolbarModule,
	MatExpansionModule,
	MatListModule,

    AngularSplitModule,
  ],
  exports: [  
	  
	
	ClassdiagramsTableComponent,
	ClassdiagramAdderComponent,
	ClassdiagramDetailComponent,
	ClassdiagramPresentationComponent,

	ClassshapesTableComponent,
	ClassshapeAdderComponent,
	ClassshapeDetailComponent,
	ClassshapePresentationComponent,

	FieldsTableComponent,
	FieldAdderComponent,
	FieldDetailComponent,
	FieldPresentationComponent,

	GorgoactionsTableComponent,
	GorgoactionAdderComponent,
	GorgoactionDetailComponent,
	GorgoactionPresentationComponent,

	LinksTableComponent,
	LinkAdderComponent,
	LinkDetailComponent,
	LinkPresentationComponent,

	PkgeltsTableComponent,
	PkgeltAdderComponent,
	PkgeltDetailComponent,
	PkgeltPresentationComponent,

	PositionsTableComponent,
	PositionAdderComponent,
	PositionDetailComponent,
	PositionPresentationComponent,

	StatesTableComponent,
	StateAdderComponent,
	StateDetailComponent,
	StatePresentationComponent,

	UmlscsTableComponent,
	UmlscAdderComponent,
	UmlscDetailComponent,
	UmlscPresentationComponent,

	VerticesTableComponent,
	VerticeAdderComponent,
	VerticeDetailComponent,
	VerticePresentationComponent,


  	SplitterComponent,
  	SidebarComponent,

  ]
})
export class GorgoModule { }
