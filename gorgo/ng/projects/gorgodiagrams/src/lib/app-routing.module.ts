import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { ClassDiagramComponent } from './class-diagram/class-diagram.component'
import { UmlscDiagramComponent } from './umlsc-diagram/umlsc-diagram.component'

const routes: Routes = [

  { path: 'classdiagram-detail/:id', component: ClassDiagramComponent, outlet: 'diagrameditor'},
  { path: 'umlsc-detail/:id', component: UmlscDiagramComponent, outlet: 'diagrameditor' },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
