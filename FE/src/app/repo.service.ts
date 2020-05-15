import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable, throwError } from 'rxjs';
import { first, catchError } from 'rxjs/operators';
import TestedRepoModel, { Repo } from './repo.model';

const API_URL: string = 'http://localhost:8080/info';

@Injectable({
  providedIn: 'root',
})
export class RepoService {
  constructor(private http: HttpClient) {}

  getRepo(payload: Repo): Observable<TestedRepoModel> {
    return this.http.post<TestedRepoModel>(API_URL, JSON.stringify(payload));
  }
}
