import { NgModule } from '@angular/core';

import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { HttpClientModule } from '@angular/common/http';
import { AppRoutingModule } from './app-routing.module'

import { GorgoModule } from 'gorgo'

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

import { AngularSplitModule } from 'angular-split';

import { GorgodiagramsComponent } from './gorgodiagrams.component';
import { UmlscDiagramComponent } from './umlsc-diagram/umlsc-diagram.component';
import { ClassDiagramComponent } from './class-diagram/class-diagram.component';
import { PkgeltDocsComponent } from './pkgelt-docs/pkgelt-docs.component';
import { ClassdiagramsSimpleTableComponent } from './classdiagrams-simple-table/classdiagrams-simple-table.component'
import { UmlscSimpleTableComponent } from './umlsc-simple/umlsc-simple.component';
import { ActionButtonsComponent } from './action-buttons/action-buttons.component';

@NgModule({
  declarations: [
    GorgodiagramsComponent, 
    UmlscDiagramComponent, 
    ClassDiagramComponent, 
    PkgeltDocsComponent,
    ClassdiagramsSimpleTableComponent,
    UmlscSimpleTableComponent,    
    ActionButtonsComponent
  ],
  imports: [

    BrowserAnimationsModule,
    HttpClientModule,
    AppRoutingModule,


    AngularSplitModule,

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

    GorgoModule

  ],
  exports: [
    PkgeltDocsComponent,
    UmlscDiagramComponent,
  ]
})
export class GorgodiagramsModule { }
