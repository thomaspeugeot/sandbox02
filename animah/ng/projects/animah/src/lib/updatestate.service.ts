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

import { UpdateStateAPI } from './updatestate-api';
import { UpdateStateDB } from './updatestate-db';



@Injectable({
  providedIn: 'root'
})
export class UpdateStateService {

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  UpdateStateServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private updatestatesUrl = 'http://localhost:8080/updatestates';

  constructor(
    private http: HttpClient
  ) { }

  /** GET updatestates from the server */
  getUpdateStates(): Observable<UpdateStateDB[]> {
    return this.http.get<UpdateStateDB[]>(this.updatestatesUrl)
      .pipe(
        tap(_ => this.log('fetched updatestates')),
        catchError(this.handleError<UpdateStateDB[]>('getUpdateStates', []))
      );
  }

  /** GET updatestate by id. Will 404 if id not found */
  getUpdateState(id: number): Observable<UpdateStateDB> {
    const url = `${this.updatestatesUrl}/${id}`;
    return this.http.get<UpdateStateDB>(url).pipe(
      tap(_ => this.log(`fetched updatestate id=${id}`)),
      catchError(this.handleError<UpdateStateDB>(`getUpdateState id=${id}`))
    );
  }

  //////// Save methods //////////

  /** POST: add a new updatestate to the server */
  postUpdateState(updatestateAPI: UpdateStateAPI): Observable<UpdateStateDB> {
    return this.http.post<UpdateStateDB>(this.updatestatesUrl, updatestateAPI, this.httpOptions).pipe(
      tap((newUpdateState: UpdateStateDB) => this.log(`added updatestate w/ id=${newUpdateState.ID}`)),
      catchError(this.handleError<UpdateStateDB>('addUpdateState'))
    );
  }

  /** DELETE: delete the updatestatedb from the server */
  deleteUpdateState(updatestatedb: UpdateStateDB | number): Observable<UpdateStateDB> {
    const id = typeof updatestatedb === 'number' ? updatestatedb : updatestatedb.ID;
    const url = `${this.updatestatesUrl}/${id}`;

    return this.http.delete<UpdateStateDB>(url, this.httpOptions).pipe(
      tap(_ => this.log(`deleted updatestatedb id=${id}`)),
      catchError(this.handleError<UpdateStateDB>('deleteUpdateState'))
    );
  }

  /** PUT: update the updatestatedb on the server */
  updateUpdateState(updatestatedb: UpdateStateDB): Observable<UpdateStateDB> {
    const id = typeof updatestatedb === 'number' ? updatestatedb : updatestatedb.ID;
    const url = `${this.updatestatesUrl}/${id}`;

    return this.http.put(url, updatestatedb, this.httpOptions).pipe(
      tap(_ => this.log(`updated updatestatedb id=${updatestatedb.ID}`)),
      catchError(this.handleError<UpdateStateDB>('updateUpdateState'))
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
