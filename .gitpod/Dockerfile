FROM gitpod/workspace-full

# Install task and gomo
RUN brew install go-task/tap/go-task abdfnx/tap/gomo

# Install packages
RUN gomo

# Install tagger
RUN curl -sL https://abdfnx.github.io/tagger/install.sh | bash
