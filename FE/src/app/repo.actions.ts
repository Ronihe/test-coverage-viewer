import { createAction, props } from '@ngrx/store';

export const loadRepos = createAction(
  '[Repo] Load Repos'
);

export const loadReposSuccess = createAction(
  '[Repo] Load Repos Success',
  props<{ data: any }>()
);

export const loadReposFailure = createAction(
  '[Repo] Load Repos Failure',
  props<{ error: any }>()
);
