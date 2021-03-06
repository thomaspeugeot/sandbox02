 // generated by GenNgService.go
import { Injectable } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';
import { HttpClient, HttpHeaders } from '@angular/common/http';

/*
 * Kamar : Ajout d'un moyen pour communiquer entre les composants qui partagent de l'information
 * afin qu'ils soient notifiés d'un changement.
 */
import { BehaviorSubject } from 'rxjs';
import { Observable, of } from 'rxjs';
import { catchError, map, tap } from 'rxjs/operators';

import { ParagraphAPI } from './paragraph-api';
import { ParagraphDB } from './paragraph-db';



@Injectable({
  providedIn: 'root'
})
export class ParagraphService {

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  ParagraphServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private paragraphsUrl = 'http://localhost:8080/paragraphs';

  constructor(
    private http: HttpClient
  ) { }

  /** GET paragraphs from the server */
  getParagraphs(): Observable<ParagraphDB[]> {
    return this.http.get<ParagraphDB[]>(this.paragraphsUrl)
      .pipe(
        tap(_ => this.log('fetched paragraphs')),
        catchError(this.handleError<ParagraphDB[]>('getParagraphs', []))
      );
  }

  /** GET paragraph by id. Will 404 if id not found */
  getParagraph(id: number): Observable<ParagraphDB> {
    const url = `${this.paragraphsUrl}/${id}`;
    return this.http.get<ParagraphDB>(url).pipe(
      tap(_ => this.log(`fetched paragraph id=${id}`)),
      catchError(this.handleError<ParagraphDB>(`getParagraph id=${id}`))
    );
  }

  //////// Save methods //////////

  /** POST: add a new paragraph to the server */
  postParagraph(paragraphAPI: ParagraphAPI): Observable<ParagraphDB> {
    return this.http.post<ParagraphDB>(this.paragraphsUrl, paragraphAPI, this.httpOptions).pipe(
      tap((newParagraph: ParagraphDB) => this.log(`added paragraph w/ id=${newParagraph.ID}`)),
      catchError(this.handleError<ParagraphDB>('addParagraph'))
    );
  }

  /** DELETE: delete the paragraphdb from the server */
  deleteParagraph(paragraphdb: ParagraphDB | number): Observable<ParagraphDB> {
    const id = typeof paragraphdb === 'number' ? paragraphdb : paragraphdb.ID;
    const url = `${this.paragraphsUrl}/${id}`;

    return this.http.delete<ParagraphDB>(url, this.httpOptions).pipe(
      tap(_ => this.log(`deleted paragraphdb id=${id}`)),
      catchError(this.handleError<ParagraphDB>('deleteParagraph'))
    );
  }

  /** PUT: update the paragraphdb on the server */
  updateParagraph(paragraphdb: ParagraphDB): Observable<ParagraphDB> {
    const id = typeof paragraphdb === 'number' ? paragraphdb : paragraphdb.ID;
    const url = `${this.paragraphsUrl}/${id}`;

    return this.http.put(url, paragraphdb, this.httpOptions).pipe(
      tap(_ => this.log(`updated paragraphdb id=${paragraphdb.ID}`)),
      catchError(this.handleError<ParagraphDB>('updateParagraph'))
    );
  }



  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error(error); // log to console instead

      // TODO: better job of transforming error for user consumption
      this.log(`${operation} failed: ${error.message}`);

      // Let the app keep running by returning an empty result.
      return of(result as T);
    };
  }

  private log(message: string) {

  }
}
