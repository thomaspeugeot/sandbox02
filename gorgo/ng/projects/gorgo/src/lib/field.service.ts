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

import { FieldAPI } from './field-api';
import { FieldDB } from './field-db';



@Injectable({
  providedIn: 'root'
})
export class FieldService {

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  FieldServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private fieldsUrl = 'http://localhost:8080/fields';

  constructor(
    private http: HttpClient
  ) { }

  /** GET fields from the server */
  getFields(): Observable<FieldDB[]> {
    return this.http.get<FieldDB[]>(this.fieldsUrl)
      .pipe(
        tap(_ => this.log('fetched fields')),
        catchError(this.handleError<FieldDB[]>('getFields', []))
      );
  }

  /** GET field by id. Will 404 if id not found */
  getField(id: number): Observable<FieldDB> {
    const url = `${this.fieldsUrl}/${id}`;
    return this.http.get<FieldDB>(url).pipe(
      tap(_ => this.log(`fetched field id=${id}`)),
      catchError(this.handleError<FieldDB>(`getField id=${id}`))
    );
  }

  //////// Save methods //////////

  /** POST: add a new field to the server */
  postField(fieldAPI: FieldAPI): Observable<FieldDB> {
    return this.http.post<FieldDB>(this.fieldsUrl, fieldAPI, this.httpOptions).pipe(
      tap((newField: FieldDB) => this.log(`added field w/ id=${newField.ID}`)),
      catchError(this.handleError<FieldDB>('addField'))
    );
  }

  /** DELETE: delete the fielddb from the server */
  deleteField(fielddb: FieldDB | number): Observable<FieldDB> {
    const id = typeof fielddb === 'number' ? fielddb : fielddb.ID;
    const url = `${this.fieldsUrl}/${id}`;

    return this.http.delete<FieldDB>(url, this.httpOptions).pipe(
      tap(_ => this.log(`deleted fielddb id=${id}`)),
      catchError(this.handleError<FieldDB>('deleteField'))
    );
  }

  /** PUT: update the fielddb on the server */
  updateField(fielddb: FieldDB): Observable<FieldDB> {
    const id = typeof fielddb === 'number' ? fielddb : fielddb.ID;
    const url = `${this.fieldsUrl}/${id}`;

    return this.http.put(url, fielddb, this.httpOptions).pipe(
      tap(_ => this.log(`updated fielddb id=${fielddb.ID}`)),
      catchError(this.handleError<FieldDB>('updateField'))
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
