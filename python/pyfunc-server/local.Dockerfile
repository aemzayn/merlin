# Copyright 2020 The Merlin Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

FROM continuumio/miniconda3

RUN wget -qO- https://dl.google.com/dl/cloudsdk/channels/rapid/downloads/google-cloud-sdk-265.0.0-linux-x86_64.tar.gz | tar xzf -
ENV PATH=$PATH:/google-cloud-sdk/bin
COPY . .
RUN mkdir /prom_dir
ENV prometheus_multiproc_dir=/prom_dir

RUN conda env create --name model_env -f ./model/conda.yaml
RUN /bin/bash -c ". activate model_env && \
    pip install -e . && \
    python -m pyfuncserver --model_dir /model --dry_run"

CMD ["/bin/bash", "./run.sh"]
