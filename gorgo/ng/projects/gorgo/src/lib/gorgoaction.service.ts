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

import { GorgoactionAPI } from './gorgoaction-api';
import { GorgoactionDB } from './gorgoaction-db';



@Injectable({
  providedIn: 'root'
})
export class GorgoactionService {

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  GorgoactionServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private gorgoactionsUrl = 'http://localhost:8080/gorgoactions';

  constructor(
    private http: HttpClient
  ) { }

  /** GET gorgoactions from the server */
  getGorgoactions(): Observable<GorgoactionDB[]> {
    return this.http.get<GorgoactionDB[]>(this.gorgoactionsUrl)
      .pipe(
        tap(_ => this.log('fetched gorgoactions')),
        catchError(this.handleError<GorgoactionDB[]>('getGorgoactions', []))
      );
  }

  /** GET gorgoaction by id. Will 404 if id not found */
  getGorgoaction(id: number): Observable<GorgoactionDB> {
    const url = `${this.gorgoactionsUrl}/${id}`;
    return this.http.get<GorgoactionDB>(url).pipe(
      tap(_ => this.log(`fetched gorgoaction id=${id}`)),
      catchError(this.handleError<GorgoactionDB>(`getGorgoaction id=${id}`))
    );
  }

  //////// Save methods //////////

  /** POST: add a new gorgoaction to the server */
  postGorgoaction(gorgoactionAPI: GorgoactionAPI): Observable<GorgoactionDB> {
    return this.http.post<GorgoactionDB>(this.gorgoactionsUrl, gorgoactionAPI, this.httpOptions).pipe(
      tap((newGorgoaction: GorgoactionDB) => this.log(`added gorgoaction w/ id=${newGorgoaction.ID}`)),
      catchError(this.handleError<GorgoactionDB>('addGorgoaction'))
    );
  }

  /** DELETE: delete the gorgoactiondb from the server */
  deleteGorgoaction(gorgoactiondb: GorgoactionDB | number): Observable<GorgoactionDB> {
    const id = typeof gorgoactiondb === 'number' ? gorgoactiondb : gorgoactiondb.ID;
    const url = `${this.gorgoactionsUrl}/${id}`;

    return this.http.delete<GorgoactionDB>(url, this.httpOptions).pipe(
      tap(_ => this.log(`deleted gorgoactiondb id=${id}`)),
      catchError(this.handleError<GorgoactionDB>('deleteGorgoaction'))
    );
  }

  /** PUT: update the gorgoactiondb on the server */
  updateGorgoaction(gorgoactiondb: GorgoactionDB): Observable<GorgoactionDB> {
    const id = typeof gorgoactiondb === 'number' ? gorgoactiondb : gorgoactiondb.ID;
    const url = `${this.gorgoactionsUrl}/${id}`;

    return this.http.put(url, gorgoactiondb, this.httpOptions).pipe(
      tap(_ => this.log(`updated gorgoactiondb id=${gorgoactiondb.ID}`)),
      catchError(this.handleError<GorgoactionDB>('updateGorgoaction'))
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
