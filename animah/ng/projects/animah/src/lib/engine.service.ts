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

import { EngineAPI } from './engine-api';
import { EngineDB } from './engine-db';


// import of struct with pointer field to current struct
import { AgentDB} from './agent-db'

@Injectable({
  providedIn: 'root'
})
export class EngineService {

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  EngineServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private enginesUrl = 'http://localhost:8080/engines';

  constructor(
    private http: HttpClient
  ) { }

  /** GET engines from the server */
  getEngines(): Observable<EngineDB[]> {
    return this.http.get<EngineDB[]>(this.enginesUrl)
      .pipe(
        tap(_ => this.log('fetched engines')),
        catchError(this.handleError<EngineDB[]>('getEngines', []))
      );
  }

  /** GET engine by id. Will 404 if id not found */
  getEngine(id: number): Observable<EngineDB> {
    const url = `${this.enginesUrl}/${id}`;
    return this.http.get<EngineDB>(url).pipe(
      tap(_ => this.log(`fetched engine id=${id}`)),
      catchError(this.handleError<EngineDB>(`getEngine id=${id}`))
    );
  }

  //////// Save methods //////////

  /** POST: add a new engine to the server */
  postEngine(engineAPI: EngineAPI): Observable<EngineDB> {
    return this.http.post<EngineDB>(this.enginesUrl, engineAPI, this.httpOptions).pipe(
      tap((newEngine: EngineDB) => this.log(`added engine w/ id=${newEngine.ID}`)),
      catchError(this.handleError<EngineDB>('addEngine'))
    );
  }

  /** DELETE: delete the enginedb from the server */
  deleteEngine(enginedb: EngineDB | number): Observable<EngineDB> {
    const id = typeof enginedb === 'number' ? enginedb : enginedb.ID;
    const url = `${this.enginesUrl}/${id}`;

    return this.http.delete<EngineDB>(url, this.httpOptions).pipe(
      tap(_ => this.log(`deleted enginedb id=${id}`)),
      catchError(this.handleError<EngineDB>('deleteEngine'))
    );
  }

  /** PUT: update the enginedb on the server */
  updateEngine(enginedb: EngineDB): Observable<EngineDB> {
    const id = typeof enginedb === 'number' ? enginedb : enginedb.ID;
    const url = `${this.enginesUrl}/${id}`;

    return this.http.put(url, enginedb, this.httpOptions).pipe(
      tap(_ => this.log(`updated enginedb id=${enginedb.ID}`)),
      catchError(this.handleError<EngineDB>('updateEngine'))
    );
  }


    // getter of struct with pointer to current struct
    getEngineAgentsViaEngine(id: number): Observable<Array<AgentDB>> {
      const url = `${this.enginesUrl}/${id}/agentsviaengine`;
      return this.http.get<Array<AgentDB>>(url).pipe(
        tap(_ => this.log(`fetched engine id=${id}`)),
        catchError(this.handleError<Array<AgentDB>>(`getEngine id=${id}`))
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
