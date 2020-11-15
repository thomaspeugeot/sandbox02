import { Component, OnInit } from '@angular/core';
import { Router, RouterState } from '@angular/router';

@Component({
  selector: 'app-sidebar',
  templateUrl: './sidebar.component.html',
  styleUrls: ['./sidebar.component.css']
})
export class SidebarComponent implements OnInit {

  constructor(
	private router: Router,
  ) { }

  ngOnInit(): void {
  }

  setTableRouterOutlet(path) {
    this.router.navigate([{
      outlets: {
        table: [path]
      }
    }]);
  }
  
  setEditorRouterOutlet(path) {
    this.router.navigate([{
      outlets: {
        editor: [path]
      }
    }]);
  }
}
