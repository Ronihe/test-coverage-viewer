import { Injectable } from '@angular/core';
import { Actions, createEffect, ofType } from '@ngrx/effects';
import { Action } from '@ngrx/store';
import { Observable, of } from 'rxjs';
import { catchError, map, mergeMap } from 'rxjs/operators';
import * as RepoActions from './repo.action';
import { RepoService } from './repo.service';
import TestedRepoModel, { Repo } from './repo.model';

@Injectable()
export class RepoEffects {
  constructor(private repoService: RepoService, private action$: Actions) {}

  GetRepo$: Observable<Action> = createEffect(() => {
    console.log('what happened in effect--');
    return this.action$.pipe(
      ofType(RepoActions.GetRepoAction),
      mergeMap((action) =>
        this.repoService.getRepo(action.payload).pipe(
          map((data: TestedRepoModel) => {
            return RepoActions.SuccessGetToDoAction({ payload: data });
          }),
          catchError((error: string) => {
            return of(
              RepoActions.ErrorRepoAction({ payload: 'error getting repo' })
            );
          })
        )
      )
    );
  });
}
