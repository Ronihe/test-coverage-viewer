import { createAction, props } from '@ngrx/store';
import TestedRepoModel, { Repo } from './repo.model';

export const GetRepoAction = createAction(
  '[Repo] - Get Repo',
  props<{ payload: Repo }>()
);

export const SuccessGetToDoAction = createAction(
  '[Repo] - Sucess Get Repo',
  props<{ payload: TestedRepoModel }>()
);

export const ErrorRepoAction = createAction(
  '[Repo] - Error',
  props<{ payload: string }>()
);
