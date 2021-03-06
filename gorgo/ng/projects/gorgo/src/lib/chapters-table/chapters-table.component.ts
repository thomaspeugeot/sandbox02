// generated by genNgTable.go
import { Component, OnInit, OnChanges, Input, Output, EventEmitter } from '@angular/core';
import { MatTableDataSource } from '@angular/material/table';
import { MatButton } from '@angular/material/button'

import { Router, RouterState } from '@angular/router';
import { ChapterDB } from '../chapter-db'
import { ChapterService } from '../chapter.service'


// generated table component
@Component({
  selector: 'app-chapters-table',
  templateUrl: './chapters-table.component.html',
  styleUrls: ['./chapters-table.component.css']
})
export class ChaptersTableComponent implements OnInit {

  // the data source for the table
  chapters: ChapterDB[];

  @Input() ID : number; // ID of the caller when component called from struct in reverse relation
  @Input() struct : string; // struct with pointer to Chapter
  @Input() field : string; // field to display

  displayedColumns: string[] = ['ID', 'Name', 'Title', 'Edit', 'Delete'];

  constructor(
    private chapterService: ChapterService,

    private router: Router,
  ) {
    // observable for changes in structs
    this.chapterService.ChapterServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getChapters()
        }
      }
    )
  }

  ngOnInit(): void {
    this.getChapters()
  }

  getChapters(): void {
    if (this.ID == null) {
      this.chapterService.getChapters().subscribe(
        Chapters => {
          this.chapters = Chapters;
        }
      )
    }
  
  }

  // newChapter initiate a new chapter
  // create a new Chapter objet
  newChapter() {
  }

  deleteChapter(chapterID: number, chapter: ChapterDB) {
    // la liste des chapters est amputée du chapter avant le delete afin
    // de mettre à jour l'IHM
    this.chapters = this.chapters.filter(h => h !== chapter);

    this.chapterService.deleteChapter(chapterID).subscribe();
  }

  editChapter(chapterID: number, chapter: ChapterDB) {

  }

  // display chapter in router
  displayChapterInRouter(chapterID: number) {
    this.router.navigate( ["chapter-display", chapterID])
  }

  // set editor outlet
  setEditorRouterOutlet(chapterID: number) {
    this.router.navigate([{
      outlets: {
        editor: ["chapter-detail", chapterID]
      }
    }]);
  }

  // set presentation outlet
  setPresentationRouterOutlet(chapterID: number) {
    this.router.navigate([{
      outlets: {
        presentation: ["chapter-presentation", chapterID]
      }
    }]);
  }
}
