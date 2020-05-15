import RepoState, { initializeState } from './repo.state';
import * as RepoActions from './repo.action';
import { createReducer, on, Action } from '@ngrx/store';
import TestedRepoModel from './repo.model';

const initialState = initializeState();

const reducer = createReducer(
  initialState,
  on(RepoActions.GetRepoAction, (state: RepoState) => {
    return {
      ...state,
      loading: true,
      loaded: false,
    };
  }),
  on(RepoActions.SuccessGetToDoAction, (state: RepoState, { payload }) => {
    return {
      ...state,
      loading: false,
      loaded: true,
      error: undefined,
      tested: payload,
    };
  }),
  on(RepoActions.ErrorRepoAction, (state: RepoState, { payload }) => {
    return {
      ...state,
      loading: false,
      loaded: false,
      error: payload,
      tested: undefined,
    };
  })
);

export function RepoReducer(
  state: RepoState | undefined,
  action: Action
): RepoState {
  return reducer(state, action);
}
