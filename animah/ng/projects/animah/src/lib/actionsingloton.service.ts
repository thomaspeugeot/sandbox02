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

import { ActionSinglotonAPI } from './actionsingloton-api';
import { ActionSinglotonDB } from './actionsingloton-db';



@Injectable({
  providedIn: 'root'
})
export class ActionSinglotonService {

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  ActionSinglotonServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private actionsinglotonsUrl = 'http://localhost:8080/actionsinglotons';

  constructor(
    private http: HttpClient
  ) { }

  /** GET actionsinglotons from the server */
  getActionSinglotons(): Observable<ActionSinglotonDB[]> {
    return this.http.get<ActionSinglotonDB[]>(this.actionsinglotonsUrl)
      .pipe(
        tap(_ => this.log('fetched actionsinglotons')),
        catchError(this.handleError<ActionSinglotonDB[]>('getActionSinglotons', []))
      );
  }

  /** GET actionsingloton by id. Will 404 if id not found */
  getActionSingloton(id: number): Observable<ActionSinglotonDB> {
    const url = `${this.actionsinglotonsUrl}/${id}`;
    return this.http.get<ActionSinglotonDB>(url).pipe(
      tap(_ => this.log(`fetched actionsingloton id=${id}`)),
      catchError(this.handleError<ActionSinglotonDB>(`getActionSingloton id=${id}`))
    );
  }

  //////// Save methods //////////

  /** POST: add a new actionsingloton to the server */
  postActionSingloton(actionsinglotonAPI: ActionSinglotonAPI): Observable<ActionSinglotonDB> {
    return this.http.post<ActionSinglotonDB>(this.actionsinglotonsUrl, actionsinglotonAPI, this.httpOptions).pipe(
      tap((newActionSingloton: ActionSinglotonDB) => this.log(`added actionsingloton w/ id=${newActionSingloton.ID}`)),
      catchError(this.handleError<ActionSinglotonDB>('addActionSingloton'))
    );
  }

  /** DELETE: delete the actionsinglotondb from the server */
  deleteActionSingloton(actionsinglotondb: ActionSinglotonDB | number): Observable<ActionSinglotonDB> {
    const id = typeof actionsinglotondb === 'number' ? actionsinglotondb : actionsinglotondb.ID;
    const url = `${this.actionsinglotonsUrl}/${id}`;

    return this.http.delete<ActionSinglotonDB>(url, this.httpOptions).pipe(
      tap(_ => this.log(`deleted actionsinglotondb id=${id}`)),
      catchError(this.handleError<ActionSinglotonDB>('deleteActionSingloton'))
    );
  }

  /** PUT: update the actionsinglotondb on the server */
  updateActionSingloton(actionsinglotondb: ActionSinglotonDB): Observable<ActionSinglotonDB> {
    const id = typeof actionsinglotondb === 'number' ? actionsinglotondb : actionsinglotondb.ID;
    const url = `${this.actionsinglotonsUrl}/${id}`;

    return this.http.put(url, actionsinglotondb, this.httpOptions).pipe(
      tap(_ => this.log(`updated actionsinglotondb id=${actionsinglotondb.ID}`)),
      catchError(this.handleError<ActionSinglotonDB>('updateActionSingloton'))
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
