import { Component, OnInit } from '@angular/core';
import { Store, select } from '@ngrx/store';

import { Repo } from '../repo.model';
import RepoState from '../repo.state';
import { getTestBed } from '@angular/core/testing';
import * as RepoActions from '../repo.action';

@Component({
  selector: 'repo',
  templateUrl: './repo.component.html',
  styleUrls: ['./repo.component.scss'],
})
export class RepoComponent implements OnInit {
  constructor(private store: Store<RepoState>) {
    //this.repo$ = store.pipe(select('tested'));
  }

  ngOnInit(): void {}

  getTested() {
    // const owner = this.userInput;
    // const repoName =
    const repo: Repo = { owner: 'google', repoName: 'uuid' };
    this.store.dispatch(RepoActions.GetRepoAction({ payload: repo }));
  }
}
