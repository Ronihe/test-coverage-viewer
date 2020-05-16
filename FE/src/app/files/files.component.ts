import { Component, OnInit } from '@angular/core';
import { Store, select } from '@ngrx/store';
import { Observable } from 'rxjs';
import { tap } from 'rxjs/operators';
import { Router } from '@angular/router';

import { Repo, FileModel } from '../repo.model';
import RepoState from '../repo.state';
import { getTestBed } from '@angular/core/testing';
import * as RepoActions from '../repo.action';
import TestedRepoModel from '../repo.model';
import { THIS_EXPR } from '@angular/compiler/src/output/output_ast';

@Component({
  selector: 'files',
  templateUrl: './files.component.html',
  styleUrls: ['./files.component.scss'],
})
export class FilesComponent implements OnInit {
  repoSub$: Observable<RepoState>;
  repoObject: RepoState;
  repo: TestedRepoModel;
  currentFile: FileModel;
  constructor(
    private store: Store<{ repo: RepoState }>,
    private route: Router
  ) {
    this.repoSub$ = this.store.select('repo').pipe(
      tap((result) => {
        this.repoObject = result;
        console.log('what is the repo obk', this.repoObject);
        this.repo = this.repoObject.tested;
        console.log('what is the repo', this.repo);
        if (this.repo) {
          this.currentFile = this.repo.goFiles[0];
        }
      })
    );
  }

  changeFile(FileId) {
    console.log(FileId);
    // this.currentFile =
  }

  ngOnInit(): void {}
}
