// Copyright 2018 Google LLC All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package remote provides facilities for reading/writing v1.Images from/to
// a remote image registry.
//
// This package supports resumable downloads via HTTP range requests. Use
// LayerRange to download specific byte ranges of layer blobs, which is useful
// for resuming interrupted downloads or implementing progressive loading.
package remote
