import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';


import { ActionsTableComponent } from './actions-table/actions-table.component'
import { ActionAdderComponent } from './action-adder/action-adder.component'
import { ActionDetailComponent } from './action-detail/action-detail.component'
import { ActionPresentationComponent } from './action-presentation/action-presentation.component'

import { AgentsTableComponent } from './agents-table/agents-table.component'
import { AgentAdderComponent } from './agent-adder/agent-adder.component'
import { AgentDetailComponent } from './agent-detail/agent-detail.component'
import { AgentPresentationComponent } from './agent-presentation/agent-presentation.component'

import { EnginesTableComponent } from './engines-table/engines-table.component'
import { EngineAdderComponent } from './engine-adder/engine-adder.component'
import { EngineDetailComponent } from './engine-detail/engine-detail.component'
import { EnginePresentationComponent } from './engine-presentation/engine-presentation.component'

const routes: Routes = [
	{ path: 'actions', component: ActionsTableComponent, outlet: 'table' },
	{ path: 'action-adder', component: ActionAdderComponent, outlet: 'editor' },
	{ path: 'action-detail/:id', component: ActionDetailComponent, outlet: 'editor' },
	{ path: 'action-presentation/:id', component: ActionPresentationComponent, outlet: 'presentation' },
	{ path: 'action-presentation-special/:id', component: ActionPresentationComponent, outlet: 'actionpres' },

	{ path: 'agents', component: AgentsTableComponent, outlet: 'table' },
	{ path: 'agent-adder', component: AgentAdderComponent, outlet: 'editor' },
	{ path: 'agent-detail/:id', component: AgentDetailComponent, outlet: 'editor' },
	{ path: 'agent-presentation/:id', component: AgentPresentationComponent, outlet: 'presentation' },
	{ path: 'agent-presentation-special/:id', component: AgentPresentationComponent, outlet: 'agentpres' },

	{ path: 'engines', component: EnginesTableComponent, outlet: 'table' },
	{ path: 'engine-adder', component: EngineAdderComponent, outlet: 'editor' },
	{ path: 'engine-detail/:id', component: EngineDetailComponent, outlet: 'editor' },
	{ path: 'engine-presentation/:id', component: EnginePresentationComponent, outlet: 'presentation' },
	{ path: 'engine-presentation-special/:id', component: EnginePresentationComponent, outlet: 'enginepres' },

];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
